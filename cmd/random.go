package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Random(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 0:
		fmt.Println("opening a random item's url in browser")
	case 1:
		fmt.Printf("open a random item's url for list %s in browser\n", args[0])
	default:
		msg := "Usage:..."
		log.Fatal(msg)
	}
}
