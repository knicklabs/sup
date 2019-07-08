package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/utils/cmd"
	"github.com/knicklabs/sup/utils"
)

// Copy copies Yesterday's and Today's tasks to the
// clipboard.
func Copy(c *cli.Context) error {
	txt, err := utils.TodayAndPreviousTasks()
	if err != nil {
		return err
	}
	cmd.Copy(txt)
	fmt.Println("Tasks copied to clipboard")
	return nil
}
