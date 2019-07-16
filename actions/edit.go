package actions

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"
)

// Edit opens today's tasks in default editor.
func Edit(c *cli.Context) error {
	var editor string

	cfg, err := config.Get()
	if err != nil {
		return err
	}

	col, err := tasks.NewCollection(cfg)
	if err != nil {
		return err
	}

	if len(c.String("editor")) > 0 {
		editor = c.String("editor")
	} else {
		editor = cfg.Editor
	}

	if c.Bool("prev") == true {
		cmd.Open(col.PrevFile, editor)
	} else {
		cmd.Open(col.CurrFile, editor)
	}
	
	return nil
}
