// Package main is the entry point of the application
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "how to use urfave cli"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "pass in some arguments to see what happens",
			Action:  test,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}

func test(c *cli.Context) error {
	// the only arguments that get passed in are the ones that come after the command
	fmt.Println(c.Args())
	return nil
}
