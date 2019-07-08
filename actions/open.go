package actions

import (
	"github.com/knicklabs/sup/utils/cmd"
	"github.com/knicklabs/sup/utils"

	"gopkg.in/urfave/cli.v1"
)

// Open opens the task directory
func Open(c *cli.Context) error {
	dir, err := utils.GetTasksDir()
	if err != nil {
		return err
	}

	cmd.Open(dir, "default")
	return nil
}
