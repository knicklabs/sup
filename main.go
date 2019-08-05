package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/knicklabs/sup/commands"
)

func main() {
	app := cli.NewApp()

	app.Name = "sup"
	app.Usage = "simple task tracker for daily standups"
	app.Version = "0.3.0"
	app.Commands = []cli.Command{
		commands.Add,
		commands.Config,
		commands.Copy,
		commands.List,
		commands.Edit,
		commands.Print,
		commands.Open,
		commands.Which,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
