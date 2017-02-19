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
			Usage:       "The IP or debug name of the Xbox",
			Destination: &consoleIPAddress,
		},
		cli.StringFlag{
			Name:        "log-level, l",
			Usage:       "The level of logging required",
			Value:       "error",
			Destination: &logLevel,
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "explorer",
			Usage:  "Explore the Xbox's file system",
			Action: commands.Explore,
		}, // explorer

		cli.Command{
			Name:   "freeze",
			Usage:  "Freeze the Xbox's threads",
			Action: commands.Freeze,
		}, // freeze
		cli.Command{
			Name:   "unfreeze",
			Usage:  "Unfreeze the Xbox's threads",
			Action: commands.Unfreeze,
		}, // unfreeze

		cli.Command{
			Name:  "memset",
			Usage: "Write memory to the Xbox",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "address, a",
					Usage: "The offset in the Xbox's memory to set",
				},
				cli.StringFlag{
					Name:  "data, d",
					Usage: "The data to write to the Xbox's memory. Must be a hex formatted byte array",
				},
			},
			Action: commands.SetMemory,
		}, // memset
		cli.Command{
			Name:  "getmem",
			Usage: "Read memory from the Xbox",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "address, a",
					Usage: "The offset in the Xbox's memory to get.",
				},
				cli.Int64Flag{
					Name:  "length, l",
					Usage: "The length of the data to read from the Xbox's memory.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "The file to write the read memory to.",
				},
			},
			Action: commands.GetMemory,
		}, // getmem
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
