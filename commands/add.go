package commands

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/actions"
)

// Add command
var Add = cli.Command{
	Name:      "add",
	ShortName: "a",
	Aliases:   []string{"new", "n"},
	Usage:     "Add a new task for Today",
	Action:    actions.Add,
}
