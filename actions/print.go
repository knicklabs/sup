package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"
)

// Print prints Yesterday's and Today's tasks.
func Print(c *cli.Context) error {
	txt, err := tasks.CurrentAndPrevious()
	if err != nil {
		return err
	}
	fmt.Println(txt)
	if c.Bool("copy") == true {
		cmd.Copy(txt)
	}
	return nil
}
