package actions

import (
	"fmt"
	
	"github.com/knicklabs/sup/utils"

	"gopkg.in/urfave/cli.v1"
)

// Which displays the location of tasks
func Which(c *cli.Context) error {
	dir, err := utils.GetTasksDir()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}
