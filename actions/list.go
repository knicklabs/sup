package actions

import (
	"fmt"

	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"

	"gopkg.in/urfave/cli.v1"
)

// List lists today's tasks or lists the previous
// day's tasks if `prev` flag is true.
func List(c *cli.Context) error {
	var txt string
	var err error
	
	if c.Bool("prev") == true {
		txt, err = tasks.Previous()
	} else {
		txt, err = tasks.Current()
	}
	
	if err != nil {
		return err
	}

	fmt.Println(txt)
	if c.Bool("copy") == true {
		cmd.Copy(txt)
	}

	return nil
}
