package actions

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"
)

// Copy copies Yesterday's and Today's tasks to the clipboard.
func Copy(c *cli.Context) error {
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	col, err := tasks.NewCollection(cfg)
	if err != nil {
		return err
	}

	dat, err := col.CurrentAndPrevious()
	if err != nil {
		return err
	}

	cmd.Copy(dat)
	return nil
}
