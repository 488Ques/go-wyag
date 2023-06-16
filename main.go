package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}

	app.Commands = []*cli.Command{
		{
			Name:  "init",
			Usage: "Initialize a new repository",
			Action: func(cCtx *cli.Context) error {
				fmt.Println("TODO Implement wyag init")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
