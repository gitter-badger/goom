package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/goom/cmd"
)

const (
	Version = "0.0.1"

	descAll    = "show all items in all lists //TODO"
	descEdit   = "edit the boom JSON file in $EDITOR //TODO"
	descDelete = "deletes a list //TODO"
	descOpen   = "open item's url in browser //TODO"
	descRandom = "open a random item's url in browser //TODO"
	descEcho   = "echo the item's value without copying //TODO"
	descCopy   = "copy the item's value without echo //TODO"
)

func main() {
	app := cli.NewApp()
	app.Name = "goom"
	app.Usage = "goom manages your text snippets on your command line like Zach Holman'boom"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:   "all",
			Usage:  descAll,
			Action: cmd.All,
		},
		{
			Name:   "edit",
			Usage:  descEdit,
			Action: cmd.Edit,
		},
		{
			Name:   "delete",
			Usage:  descDelete,
			Action: cmd.Delete,
		},
		{
			Name:   "open",
			Usage:  descOpen,
			Action: cmd.Open,
		},
		{
			Name:   "random",
			Usage:  descRandom,
			Action: cmd.Random,
		},
		{
			Name:   "echo",
			Usage:  descEcho,
			Action: cmd.Echo,
		},
		{
			Name:   "copy",
			Usage:  descCopy,
			Action: cmd.Copy,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
