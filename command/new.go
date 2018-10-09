package command

import (
	"fmt"
	"os"
	"text/template"
	"bytes"
	"bufio"
	patt "github.com/takaishi/patt/lib"
	"github.com/gernest/front"
	"path/filepath"
	"regexp"
	"time"
	"github.com/urfave/cli"
)

type Variables struct {
	Year string
	Month string
	Day string
	Week string
}

//func configFilePath() string {
//	return os.Getenv("HOME") + "/.patt.d/config.json"
//}

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

func RunNewCommand(c *cli.Context) error {
	name := c.Args().Get(0)
	configs := patt.ReadConfig()
	src := configs[name].Source

	t := time.Now()
	data := Variables{
		Year: string(t.Year()),
		Month: string(t.Month()),
		Day: string(t.Day()),
		Week: string(t.Weekday()),
	}

	doc := readTemplateFile(src, data)

	err := createFileFromTemplate(doc)
	if err != nil {
		return err
	}
	return nil
}
