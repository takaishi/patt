package template

import (
	"bytes"
	"fmt"
	"github.com/gernest/front"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

type Template struct {
	Name        string
	Source      string
	Destination string
}

type Templates []Template

type Variables struct {
	Year  string
	Month string
	Day   string
	Week  string
}

func ReadTemplates() (Templates, error) {
	templates := Templates{}
	files, err := ioutil.ReadDir(os.Getenv("HOME") + "/.patt.d/templates")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		src := os.Getenv("HOME") + "/.patt.d/templates/" + file.Name()
		doc := ReadTemplateFile(src, GetVariables())

		m := front.NewMatter()
		m.Handle("+++", front.JSONHandler)
		f, _, err := m.Parse(&doc)
		if err != nil {
			return nil, err
		}
		name := f["name"].(string)
		dst := f["destination"].(string)
		templates = append(templates, Template{Name: name, Source: src, Destination: dst})
	}
	return templates, nil
}

func ReadTemplateFile(src string, data Variables) bytes.Buffer {
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

func GetVariables() (v Variables) {
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
