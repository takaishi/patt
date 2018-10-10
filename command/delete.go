package command

import (
	"errors"
	"github.com/takaishi/patt/template"
	"github.com/urfave/cli"
	"os"
)

func RunDeleteCommand(c *cli.Context) error {
	name := c.Args().Get(0)

	templates, err := template.ReadTemplates()
	if err != nil {
		return err
	}

	for _, tmpl := range templates {
		if tmpl.Name == name {
			err := os.Remove(tmpl.Source)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("Cound not find template " + name)
}
