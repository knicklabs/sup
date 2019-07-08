package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"
)

// Copy copies Yesterday's and Today's tasks to the
// clipboard.
func Copy(c *cli.Context) error {
	txt, err := tasks.CurrentAndPrevious()
	if err != nil {
		return err
	}
	cmd.Copy(txt)
	fmt.Println("Tasks copied to clipboard")
	return nil
}
