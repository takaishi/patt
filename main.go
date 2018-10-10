package main

import (
	"github.com/takaishi/patt/command"
	"github.com/takaishi/patt/config"
	"github.com/urfave/cli"
	"log"
	"os"
)

var version string

func main() {
	app := cli.NewApp()
	app.Version = config.Version

	app.Commands = []cli.Command{
		{
			Name: "add",
			Action: func(c *cli.Context) error {
				return command.RunAddCommand(c)
			},
		},
		{
			Name: "list",
			Action: func(c *cli.Context) error {
				return command.RunListCommand(c)
			},
		},
		{
			Name: "new",
			Action: func(c *cli.Context) error {
				return command.RunNewCommand(c)
			},
		},
		{
			Name: "delete",
			Action: func(c *cli.Context) error {
				return command.RunDeleteCommand(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
