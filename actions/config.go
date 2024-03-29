package actions

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/config"
)

func output(cfg *config.Config) {
	fmt.Printf("Dir:    %s\n", cfg.Dir)
	fmt.Printf("Editor: %s\n", cfg.Editor)
}

// Config configures SUP
func Config(c *cli.Context) error {
	cfg, err := config.Get()
	if err != nil {
		return err
	}

	if c.Bool("reset") == true {
		err = cfg.Reset()
		if err != nil {
			return err
		}
		output(cfg)
		return nil
	}

	nDir := c.String("dir")
	nEdr := c.String("editor")

	isNDir := len(nDir) > 0
	isNEdr := len(nEdr) > 0

	if isNDir == true {
		cfg.Dir = nDir
	}

	if isNEdr == true {
		cfg.Editor = nEdr
	}

	if isNDir == true || isNEdr == true {
		err = cfg.Save()
		if err != nil {
			return err
		}
	}

	output(cfg)
	return nil
}
