package command

import (
	"fmt"
	"github.com/gernest/front"
	patt "github.com/takaishi/patt/lib"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
	"time"
	"github.com/urfave/cli"
	"bytes"
	"bufio"
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
	configs := patt.ReadConfig()
	src := configs[name].Source

	doc := readTemplateFile(src, getVariables())

	err := createFileFromTemplate(doc)
	if err != nil {
		return err
	}
	return nil
}
