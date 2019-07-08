package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Open command
var Open = cli.Command{
	Name:      "open",
	ShortName: "o",
	Usage:     "Opens the task directory",
	Action:    actions.Open,
}
