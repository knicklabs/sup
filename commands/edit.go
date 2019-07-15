package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Edit command
var Edit = cli.Command{
	Name:      "edit",
	ShortName: "e",
	Usage:     "Edit Today's tasks",
	Action:    actions.Edit,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "editor",
			Value: "default",
			Usage: "Choose editor for editing tasks",
		},
	},
}
