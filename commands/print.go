package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Print command
var Print = cli.Command{
	Name:      "print",
	ShortName: "p",
	Usage:     "Print Yesterday's and Today's tasks",
	Action:    actions.Print,
	Flags:     []cli.Flag{
		cli.BoolFlag{
			Name: "copy",
			Usage: "Copy the output to the clipboard",
		},
	},
}
