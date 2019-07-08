package actions

import (
	"fmt"

	"github.com/knicklabs/sup/utils"

	"gopkg.in/urfave/cli.v1"
)

// Add adds a task to Today's file.
func Add(c *cli.Context) error {
	fn := utils.GetCurrentFilename()

	dir, err := utils.GetTasksDir()
	if err != nil {
		return err
	}

	err = utils.MakeDir(dir)
	if err != nil {
		return err
	}

	err = utils.WriteStringToFile(dir, fn, fmt.Sprintf("\n- %s", c.Args().First()))
	if err != nil {
		return err
	}

	fmt.Println("Added task: ", c.Args().First())
	return nil
}
