package command

import (
	"strings"
	patt "github.com/takaishi/patt/lib"
	"fmt"
	"os"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	key := args[0]
	configs := patt.ReadConfig()

	path := configs[key].Source
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	delete(configs, key)

	err = patt.WriteConfig(configs)
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
