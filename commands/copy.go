package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Copy command
var Copy = cli.Command{
	Name:      "copy",
	ShortName: "cp",
	Usage:     "Copy Yesterday's and Today's tasks",
	Action:    actions.Copy,
}
