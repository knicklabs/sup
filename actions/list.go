package actions

import (
	"fmt"

	"github.com/knicklabs/sup/utils"

	"gopkg.in/urfave/cli.v1"
)

// List lists today's tasks or lists the previous
// day's tasks if `prev` flag is true.
func List(c *cli.Context) error {
	var txt string
	var err error
	
	if c.Bool("prev") == true {
		txt, err = utils.PreviousTasks()
	} else {
		txt, err = utils.TodaysTasks()
	}
	
	if err != nil {
		return err
	}

	fmt.Println(txt)
	return nil
}
