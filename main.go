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
			Usage: "Addd new template",
			Action: func(c *cli.Context) error {
				return command.RunAddCommand(c)
			},
		},
		{
			Name: "list",
			Usage: "Show templates",
			Action: func(c *cli.Context) error {
				return command.RunListCommand(c)
			},
		},
		{
			Name: "new",
			Usage: "Generate file from template",
			Action: func(c *cli.Context) error {
				return command.RunNewCommand(c)
			},
		},
		{
			Name: "delete",
			Usage: "Delete template",
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
