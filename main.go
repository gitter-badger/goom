package goom

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/goom/cmd"
	"github.com/spf13/viper"
)

const (
	Version = "0.0.1"

	descBase = `manages your text snippets on your command line like Zach Holman'boom
   goom <list>                   create a new list //TODO
   goom <list>                   show items for a list //TODO
   goom <list> <name> <value>    create a new list item //TODO
   goom <name>                   copy item's value to clipboard /TODO
   goom <list> <name>            copy item's value to clipboard //TODO`
	descAll    = "show all items in all lists //TODO"
	descEdit   = "edit the boom JSON file in $EDITOR //TODO"
	descDelete = "deletes a list or item from a list //TODO"
	descOpen   = "open item's url in browser //TODO"
	descRandom = "open a random item's url in browser //TODO"
	descEcho   = "echo the item's value without copying //TODO"
	descCopy   = "copy the item's value without echo //TODO"
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("goomPath", usr.HomeDir)
	viper.SetDefault("configPath", filepath.Join(viper.GetString("goomPath"), ".goomrc"))
	viper.SetDefault("dataPath", filepath.Join(viper.GetString("goomPath"), ".goom"))
	viper.SetDefault("boltdbPath", filepath.Join(viper.GetString("dataPath"), "bolt.db"))

	if _, err := os.Stat(viper.GetString("dataPath")); os.IsNotExist(err) {
		os.Mkdir(viper.GetString("dataPath"), 0750)
	}
}

func RunApp() {
	app := cli.NewApp()
	app.Name = "goom"
	app.Usage = descBase
	app.Version = Version
	app.Action = cmd.Base
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
