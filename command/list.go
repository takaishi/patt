package command

import (
	"strings"
	patt "github.com/takaishi/patt/lib"
	"fmt"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	configs := patt.ReadConfig()

	for k := range configs {
		fmt.Printf("%s %s\n", k, configs[k].Source)
	}

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
