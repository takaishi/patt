package command

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gernest/front"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"
)

type Variables struct {
	Year  string
	Month string
	Day   string
	Week  string
}

func readTemplateFile(src string, data Variables) bytes.Buffer {
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		fmt.Println(err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil {
		fmt.Println(err)
	}

	return doc
}

func createFileFromTemplate(doc bytes.Buffer) error {
	m := front.NewMatter()
	m.Handle("+++", front.JSONHandler)
	f, body, err := m.Parse(&doc)
	if err != nil {
		return err
	}

	dst := f["destination"].(string)
	base := filepath.Base(dst)
	rep := regexp.MustCompile(base + "$")
	dstDir := rep.ReplaceAllString(dst, "")
	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	fp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fp.Close()

	writer := bufio.NewWriter(fp)
	_, err = writer.WriteString(body)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func getVariables() (v Variables) {
	t := time.Now()
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	v = Variables{
		Year:  fmt.Sprintf("%d", t.Year()),
		Month: fmt.Sprintf("%02d", t.Month()),
		Day:   fmt.Sprintf("%02d", t.Day()),
		Week:  fmt.Sprintf("%s", wdays[t.Weekday()]),
	}
	return
}

func RunNewCommand(c *cli.Context) error {
	name := c.Args().Get(0)
	templates := Templates{}
	files, err := ioutil.ReadDir(os.Getenv("HOME") + "/.patt.d/templates")
	if err != nil {
		return err
	}

	for _, file := range files {
		src := os.Getenv("HOME") + "/.patt.d/templates/" + file.Name()
		doc := readTemplateFile(src, getVariables())

		m := front.NewMatter()
		m.Handle("+++", front.JSONHandler)
		f, _, err := m.Parse(&doc)
		if err != nil {
			return err
		}
		name := f["name"].(string)
		dst := f["destination"].(string)
		templates = append(templates, Template{Name: name, Source: src, Destination: dst})
	}

	for _, template := range templates {
		if template.Name == name {
			doc := readTemplateFile(template.Source, getVariables())
			err := createFileFromTemplate(doc)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Cound not find template " + name)
}
