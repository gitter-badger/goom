package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Open(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		fmt.Printf("opening %s item's url in browser\n", args[0])
	case 2:
		fmt.Printf("opening %s item's url in browser for list %s\n", args[1], args[0])
	default:
		msg := "Usage: .."
		log.Fatal(msg)
	}
}
