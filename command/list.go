package command

import (
	"github.com/olekukonko/tablewriter"
	patt "github.com/takaishi/patt/lib"
	"github.com/urfave/cli"
	"os"
)

func RunListCommand(c *cli.Context) error {
	configs := patt.ReadConfig()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Source"})
	for k := range configs {
		table.Append([]string{k, configs[k].Source})
	}

	table.Render()
	return nil
}
