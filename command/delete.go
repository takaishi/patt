package command

import (
	patt "github.com/takaishi/patt/lib"
	"github.com/urfave/cli"
	"os"
)

func RunDeleteCommand(c *cli.Context) error {
	key := c.Args().Get(0)
	configs := patt.ReadConfig()

	path := configs[key].Source
	err := os.Remove(path)
	if err != nil {
		return err
	}

	delete(configs, key)

	err = patt.WriteConfig(configs)
	if err != nil {
		return err
	}

	return nil
}
