package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Copy(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		fmt.Printf("Copying item %s value without echoing\n", args[0])
	case 2:
		fmt.Printf("Copying item %s value from list %s without echoing\n", args[1], args[0])
	default:
		msg := "Usage:..."
		log.Fatal(msg)
	}
}
