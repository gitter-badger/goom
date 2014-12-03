package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func Base(ctx *cli.Context) {
	args := ctx.Args()

	switch len(args) {
	case 1:
		fmt.Printf("Creating list %s, or showing items from list %s, or copying value of item %s to clipboard\n", args[0], args[0], args[0])
	case 2:
		fmt.Printf("Copying value of item %s from list %s to clipboard\n", args[1], args[0])
	case 3:
		fmt.Printf("Creating new item %s in list %s with value %s\n", args[1], args[0], args[2])
	default:
		msg := "Usage:...."
		log.Fatal(msg)
	}

}
