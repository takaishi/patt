package command

import (
	"strings"
)

type AddCommand struct {
	Meta
}

func (c *AddCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *AddCommand) Synopsis() string {
	return ""
}

func (c *AddCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
