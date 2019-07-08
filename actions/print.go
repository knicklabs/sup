package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/utils/cmd"
	"github.com/knicklabs/sup/utils"
)

// Print prints Yesterday's and Today's tasks.
func Print(c *cli.Context) error {
	txt, err := utils.TodayAndPreviousTasks()
	if err != nil {
		return err
	}
	fmt.Println(txt)
	if c.Bool("copy") == true {
		cmd.Copy(txt)
	}
	return nil
}
