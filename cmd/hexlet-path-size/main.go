package main

import (
	"code/cmd/code"
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

			if cmd.Args().Len() == 0 {
				return fmt.Errorf("path is required")
			}
			size, err := code.GetSize(cmd.Args().Get(0))
			if err != nil {
				return err
			}

			fmt.Printf("%d" + "B" + "\t" + "%q\n", size, cmd.Args().Get(0))
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
