package actions

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/utils/cmd"
)

// Open opens the task directory
func Open(c *cli.Context) error {
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	dir, err := cfg.AbsoluteTasksPath()
	if err != nil {
		return err
	}

	cmd.Open(dir, "default")
	return nil
}
