package actions

import (
	"fmt"

	"github.com/knicklabs/sup/config"
	"github.com/knicklabs/sup/tasks"
	"github.com/knicklabs/sup/utils/cmd"

	"gopkg.in/urfave/cli.v1"
)

// List lists today's tasks or lists the previous
// day's tasks if `prev` flag is true.
func List(c *cli.Context) error {
	var dat string
	var err error

	cfg, err := config.Get()
	if err != nil {
		return err
	}

	col, err := tasks.NewCollection(cfg)
	if err != nil {
		return err
	}

	d := c.String("date")
	
	if len(d) > 0 {
		dat, err = col.Date(d)
	} else if (c.Bool("prev") == true) {
		dat, err = col.Previous()
	} else {
		dat, err = col.Current()
	}

	if err != nil {
		return err
	}

	fmt.Println(dat)
	if c.Bool("copy") == true {
		cmd.Copy(dat)
	}

	return nil
}
