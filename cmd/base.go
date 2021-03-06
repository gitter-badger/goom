package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
	"github.com/humboldtux/goom/repo"
	"github.com/spf13/viper"
)

func Base(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		db, err := bolt.Open(viper.GetString("boltdbPath"), 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			db.Close()
		}()

		r := repo.List{db}
		err = r.Create(args[0])
		if err == nil {
			fmt.Printf("List %s created\n", args[0])
		}
	case 2:
		db, err := bolt.Open(viper.GetString("boltdbPath"), 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			db.Close()
		}()
		i := repo.Item{db, args[0]}
		v, err := i.Get(args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%s\n", v)
		}
	case 3:
		db, err := bolt.Open(viper.GetString("boltdbPath"), 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			db.Close()
		}()

		i := repo.Item{db, args[0]}
		err = i.Create(args[1], args[2])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Creating new item %s in list %s with value %s\n", args[1], args[0], args[2])
		}
	default:
		msg := "Usage:...."
		log.Fatal(msg)
	}

}
