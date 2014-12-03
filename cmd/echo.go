package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Echo(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		fmt.Printf("Echoing item %s value without copying\n", args[0])
	case 2:
		fmt.Printf("Echoing item %s value from list %s without copying\n", args[1], args[0])
	default:
		msg := "Usage:..."
		log.Fatal(msg)
	}
}
