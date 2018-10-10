package command

import (
	"bufio"
	"bytes"
	"github.com/gernest/front"
	"github.com/pkg/errors"
	"github.com/takaishi/patt/template"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"regexp"
)

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

	templates, err := template.ReadTemplates()
	if err != nil {
		return err
	}

	for _, tmpl := range templates {
		if tmpl.Name == name {
			doc := template.ReadTemplateFile(tmpl.Source, template.GetVariables())
			err := createFileFromTemplate(doc)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Cound not find template " + name)
}
