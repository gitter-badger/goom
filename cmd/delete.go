package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Delete(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		fmt.Printf("deleting list %s\n", args[0])
	case 2:
		fmt.Printf("Deleting item %s from list %s\n", args[1], args[0])
	default:
		msg := "Usage:..."
		log.Fatal(msg)
	}
}
