package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Which command
var Which = cli.Command{
	Name:      "which",
	ShortName: "w",
	Usage:     "Display location of tasks",
	Action:    actions.Which,
}
