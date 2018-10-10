package command

import (
	"github.com/olekukonko/tablewriter"
	"github.com/takaishi/patt/template"
	"github.com/urfave/cli"
	"os"
)

func RunListCommand(c *cli.Context) error {
	table := tablewriter.NewWriter(os.Stdout)
	templates, err := template.ReadTemplates()
	if err != nil {
		return err
	}

	table.SetHeader([]string{"Name", "Source"})
	for _, t := range templates {
		table.Append([]string{t.Name, t.Source})
	}

	table.Render()

	return nil
}
