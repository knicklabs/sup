package actions

import (
	"fmt"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/tasks"

	"gopkg.in/urfave/cli.v1"
)

// Add adds a task to Today's file.
func Add(c *cli.Context) error {
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	col, err := tasks.NewCollection(cfg)
	if err != nil {
		return err
	}

	return col.Add(fmt.Sprintf("\n- %s", c.Args().First()))
}
