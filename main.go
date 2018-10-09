package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/takaishi/patt/command"
	"log"
)

func main() {
	app := cli.NewApp()
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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
