package main

import (
	"os"

	"github.com/0xdeafcafe/ares/commands"

	"gopkg.in/urfave/cli.v1"
)

var (
	consoleIPAddress = ""
	logLevel         = ""
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "ip",
			Usage:       "The IP or debug name of the Xbox.",
			Destination: &consoleIPAddress,
		},
		cli.StringFlag{
			Name:        "logLevel, -l",
			Usage:       "The level of logging required",
			Value:       "error",
			Destination: &logLevel,
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "explorer",
			Usage: "Explore the Xbox's file system.",
			Action: func(c *cli.Context) error {
				explorer, err := commands.NewExplorer(c, consoleIPAddress)
				if err != nil {
					return err
				}

				explorer.StartListening()
				return explorer.Listen()
			},
		},
	}

	app.Run(os.Args)
}
