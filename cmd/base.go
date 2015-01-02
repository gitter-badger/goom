package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
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

		//NOT WORKING
		//repo := listRepo{db}
		//err = repo.create(args[0])
		if err == nil {
			fmt.Printf("List %s created", args[0])
		}
	case 2:
		fmt.Printf("Copying value of item %s from list %s to clipboard\n", args[1], args[0])
	case 3:
		fmt.Printf("Creating new item %s in list %s with value %s\n", args[1], args[0], args[2])
	default:
		msg := "Usage:...."
		log.Fatal(msg)
	}

}
