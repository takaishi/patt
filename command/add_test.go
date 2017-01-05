package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAddCommand_implement(t *testing.T) {
	var _ cli.Command = &AddCommand{}
}
