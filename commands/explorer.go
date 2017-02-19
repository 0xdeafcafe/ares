package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0xdeafcafe/ares/helpers"
	"github.com/0xdeafcafe/ares/models"
	goxbdm "github.com/0xdeafcafe/go-xbdm"
	"gopkg.in/urfave/cli.v1"
)

// Explorer ..
type Explorer struct {
	context          *cli.Context
	xbdm             *goxbdm.Client
	reader           *bufio.Reader
	currentDirectory string
	currentFolder    string
	directoryCache   map[string]*models.DirectoryCache
}

const (
	connectionEstablishedFormat = "#### Connection Established with Development Kit\n#### IP Address: %s\n#### Console Debug Name: %s"
	lineStart                   = "%s@%s:%s>"
)

func (explorer *Explorer) listen() error {
	for {
		if explorer.currentDirectory == "" {
			explorer.currentFolder = "~"
		} else {
			//explorer.currentFolder = strings.Split(explorer.currentDirectory, "\\")
		}

		fmt.Print(fmt.Sprintf(lineStart, explorer.xbdm.ConsoleName(), explorer.xbdm.ConsoleIP, explorer.currentDirectory))
		input, _ := explorer.reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")

		switch {
		case input == "exit":
			return nil
		case input == "clear":
			helpers.ClearConsole()
			break

		case input == "ls" && explorer.currentDirectory == "":
			drives, err := explorer.xbdm.ListDrives()
			if err != nil {
				return err
			}

			for _, drive := range drives {
				fmt.Println(drive.Name)
			}
			break
		case input == "ls" && explorer.currentDirectory != "":
			directoryItems, err := explorer.xbdm.ListDirectory(explorer.currentDirectory)
			if err != nil {
				return err
			}

			for _, directoryItem := range directoryItems {
				fmt.Println(directoryItem.Name)
			}
			break

		case strings.HasPrefix(input, "cd "):
			format := "cd %s"
			var path string
			fmt.Sscanf(input, format, &path)

			newCD := ""
			if explorer.currentDirectory == "" {
				newCD = fmt.Sprintf("%s:\\", path)
			} else {
				newCD = fmt.Sprintf("%s\\%s", explorer.currentDirectory, path)
			}

			directoryItems, err := explorer.xbdm.ListDirectory(newCD)
			if err != nil {
				panic(err)
			}

			explorer.directoryCache[newCD] = models.NewDirectoryCache(directoryItems)
			explorer.currentDirectory = newCD
			break
		default:
			fmt.Println("")
			break
		}
	}
}

func StartExplorer(c *cli.Context) error {
	ip := c.GlobalString("ip")
	client := &Explorer{
		context:          c,
		currentDirectory: "",
		directoryCache:   make(map[string]*models.DirectoryCache),
	}

	xbdm, err := goxbdm.NewXBDMClient(ip)
	if err != nil {
		return err
	}

	client.reader = bufio.NewReader(os.Stdin)
	client.xbdm = xbdm

	fmt.Println(fmt.Sprintf(connectionEstablishedFormat, client.xbdm.ConsoleIP, client.xbdm.ConsoleName()))
	fmt.Println()
	fmt.Println()

	return client.listen()
}
