package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0xdeafcafe/ares/helpers"
	goxbdm "github.com/0xdeafcafe/go-xbdm"
	"gopkg.in/urfave/cli.v1"
)

type Explorer struct {
	context          *cli.Context
	xbdm             *goxbdm.Client
	reader           *bufio.Reader
	currentDirectory string
}

const (
	connectionEstablishedFormat = `#### Connection Established with Development Kit
#### IP Address: %s
#### Console Debug Name: %s`
	lineStart = "%s@%s>"
)

// StartListening ..
func (explorer *Explorer) StartListening() {
	fmt.Println(fmt.Sprintf(connectionEstablishedFormat, explorer.xbdm.ConsoleIP, explorer.xbdm.ConsoleName()))
	fmt.Println()
	fmt.Println()
}

// Listen ..
func (explorer *Explorer) Listen() error {
	for {
		fmt.Print(fmt.Sprintf(lineStart, explorer.xbdm.ConsoleIP, explorer.xbdm.ConsoleName()))
		input, _ := explorer.reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		switch input {
		case "exit":
			return nil
		case "clear":
			helpers.ClearConsole()
			break

		case "ls":
			drives, err := explorer.xbdm.ListDrives()
			if err != nil {
				return err
			}

			for _, drive := range drives {
				fmt.Println(drive)
			}
			break

		case "cs":

			break
		}
	}
}

// NewExplorer ..
func NewExplorer(c *cli.Context, ip string) (*Explorer, error) {
	client := &Explorer{
		context:          c,
		currentDirectory: "/",
	}

	xbdm, err := goxbdm.NewXBDMClient(ip)
	if err != nil {
		return nil, err
	}

	client.reader = bufio.NewReader(os.Stdin)
	client.xbdm = xbdm
	return client, nil
}
