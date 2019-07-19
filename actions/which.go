package actions

import (
	"fmt"

	"github.com/knicklabs/sup/config"

	"gopkg.in/urfave/cli.v1"
)

// Which displays the location of tasks
func Which(c *cli.Context) error {
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	dir, err := cfg.AbsoluteTasksPath()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}
