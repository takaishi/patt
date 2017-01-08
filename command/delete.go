package command

import (
	"strings"
	patt "github.com/takaishi/patt/lib"
	"fmt"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	key := args[0]
	configs := patt.ReadConfig()

	delete(configs, key)

	err := patt.WriteConfig(configs)
	if err != nil {
		fmt.Println(err)
		return 1
	}


	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return ""
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
