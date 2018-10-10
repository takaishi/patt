package command

import (
	"github.com/gernest/front"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

type Template struct {
	Name        string
	Source      string
	Destination string
}

type Templates []Template

func RunListCommand(c *cli.Context) error {
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Source"})
	for _, t := range templates {
		table.Append([]string{t.Name, t.Source})
	}

	table.Render()

	return nil
}
