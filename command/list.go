package command

import (
	"strings"
	patt "github.com/takaishi/patt/lib"
	"github.com/olekukonko/tablewriter"
	"os"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	configs := patt.ReadConfig()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Source"})
	for k := range configs {
		table.Append([]string{k, configs[k].Source})
	}
	       table.Render()
	return 0
}

func (c *ListCommand) Synopsis() string {
	return ""
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
