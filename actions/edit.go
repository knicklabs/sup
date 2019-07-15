package actions

import (
	"path"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/utils"
	"github.com/knicklabs/sup/utils/cmd"
)

// Edit opens today's tasks in default editor.
func Edit(c *cli.Context) error {
	fn := utils.GetCurrentFilename()

	dir, err := utils.GetTasksDir()
	if err != nil {
		return err
	}

	cmd.Open(path.Join(dir, fn), c.String("editor"))
	return nil
}
