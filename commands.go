package main

import (
	"github.com/mitchellh/cli"
	"github.com/takaishi/patt/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &command.AddCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				Meta: *meta,
			}, nil
		},
		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
