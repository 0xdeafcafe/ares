package main

import (
	"errors"
	"os"

	"github.com/0xdeafcafe/ares/commands"
	"github.com/0xdeafcafe/go-xbdm"

	"regexp"

	"fmt"

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
		cli.Command{
			Name:  "memset",
			Usage: "Explore the Xbox's file system.",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "address, a",
					Usage: "The offset in the Xbox's memory to set.",
				},
				cli.StringFlag{
					Name:  "data, d",
					Usage: "The data to write to the Xbox's memory. Must be a hex formatted byte array.",
				},
			},
			Action: func(c *cli.Context) error {
				addr := c.Int64("address")
				data := c.String("data")

				// Validate Addr is within Range
				if addr < 0 {
					return errors.New("address can not be negative")
				}

				// Validate Data is valid hex string
				regex := regexp.MustCompile("[A-Fa-f0-9]+")
				if !regex.MatchString(data) {
					return errors.New("data invalid format")
				}
				if len(data)%2 != 0 {
					return errors.New("data invalid length")
				}

				// Create connection to xbox
				xbdm, err := goxbdm.NewXBDMClient(consoleIPAddress)
				if err != nil {
					return err
				}

				// Set memory
				resp, err := xbdm.SetMemory(addr, data)
				if err != nil {
					return err
				}

				fmt.Println(resp)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
