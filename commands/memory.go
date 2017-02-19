package commands

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"io/ioutil"

	"encoding/hex"

	"github.com/0xdeafcafe/go-xbdm"
	"gopkg.in/urfave/cli.v1"
)

func SetMemory(c *cli.Context) error {
	ip := c.GlobalString("ip")
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
	xbdm, err := goxbdm.NewXBDMClient(ip)
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
}

func GetMemory(c *cli.Context) error {
	ip := c.GlobalString("ip")
	addr := c.Int64("address")
	length := c.Int64("length")
	output := c.String("output")

	// Validate Addr is within Range
	if addr < 0 {
		return errors.New("address can not be negative")
	}

	// Validate Len is within Range
	if length < 0 {
		return errors.New("length can not be negative")
	}

	// Create connection to xbox
	xbdm, err := goxbdm.NewXBDMClient(ip)
	if err != nil {
		return err
	}

	// Get memory
	data, err := xbdm.GetMemory(addr, length)
	if err != nil {
		return err
	}

	// If we don't specify an output file, dump to term
	if output == "" {
		fmt.Println(hex.EncodeToString(data))
		return nil
	}

	return ioutil.WriteFile(output, data, os.ModePerm)
}
