package command

import (
	"errors"
	"github.com/gernest/front"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func RunDeleteCommand(c *cli.Context) error {
	templates := Templates{}
	name := c.Args().Get(0)

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
			err := os.Remove(template.Source)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("Cound not find template " + name)
}
