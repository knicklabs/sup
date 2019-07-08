package commands

import (
	"gopkg.in/urfave/cli.v1"
	
	"github.com/knicklabs/sup/actions"
)

// List command
var List = cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "List Today's tasks",
	Action:    actions.List,
	Flags:     []cli.Flag{
		cli.BoolFlag{
			Name:  "prev",
			Usage: "Show the previous day's tasks",
		},
	},
}
