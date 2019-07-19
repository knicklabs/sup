package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"
)

// Print prints Yesterday's and Today's tasks.
func Print(c *cli.Context) error {
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

	fmt.Println(dat)
	if c.Bool("copy") == true {
		cmd.Copy(dat)
	}

	return nil
}
