package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Config command
var Config = cli.Command{
	Name:   "config",
	Usage:  "configures SUP",
	Action: actions.Config,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Usage: "Customize the directory for saving tasks",
		},
		cli.StringFlag{
			Name:  "editor",
			Usage: "Customize the editor for editing tasks",
		},
		cli.BoolFlag{
			Name:  "reset",
			Usage: "Reset back to defaults",
		},
	},
}
