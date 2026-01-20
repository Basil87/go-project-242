package main

import (
	"code"
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

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "use human-readable format",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {

			if cmd.Args().Len() == 0 {
				fmt.Println("path is required")
				return nil
			}

			path := cmd.Args().Get(0)
			size, err := code.GetSize(path)
			if err != nil {
				return err
			}

			output := code.HumanSize(size, cmd.Bool("human"))
			fmt.Printf("%s\t%q\n", output, path)

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
