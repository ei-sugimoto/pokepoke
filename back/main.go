package main

import (
	"os"

	"github.com/ei-sugimoto/pokepoke/back/cmd"
	"github.com/urfave/cli"
)

func main() {
	var (
		suffix string
	)

	app := cli.NewApp()
	app.Name = "pokepoke"
	app.Usage = "pokepoke"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "suffix",
			Value:       "default",
			Usage:       "suffix",
			Destination: &suffix,
			EnvVar:      "SUFFIX",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "serve",
			Action: func(c *cli.Context) error {
				cmd.Serve()
				return nil
			},
		},
		{
			Name:  "migration",
			Usage: "migration",
			Action: func(c *cli.Context) error {
				cmd.Migration()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
