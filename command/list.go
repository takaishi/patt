package command

import (
	patt "github.com/takaishi/patt/lib"
	"fmt"
	"github.com/urfave/cli"
)

func RunListCommand(c *cli.Context) error {
	configs := patt.ReadConfig()

	for k := range configs {
		fmt.Printf("%s %s\n", k, configs[k].Source)
	}

	return nil
}