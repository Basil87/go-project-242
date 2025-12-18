package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("Hello from Hexlet!")
			return nil
		},
	}

	cli.RootCommandHelpTemplate = `
	./bin/hexlet-path-size
	
	NAME:
		hexlet-path-size - print size of a file or directory
	
	USAGE:
		hexlet-path-size [global options]
	
	GLOBAL OPTIONS:
		--help, -h  show help
`

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
