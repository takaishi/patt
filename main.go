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
			Name: "list",
			Action: func(c *cli.Context) error {
				return command.RunListCommand(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
