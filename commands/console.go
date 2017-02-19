package commands

import (
	"github.com/0xdeafcafe/go-xbdm"
	"gopkg.in/urfave/cli.v1"
)

func Freeze(c *cli.Context) (err error) {
	ip := c.GlobalString("ip")
	xbdm, err := goxbdm.NewXBDMClient(ip)
	if err != nil {
		_, err = xbdm.ChangeFreezeState(goxbdm.FrozenState)
	}

	return
}

func Unfreeze(c *cli.Context) (err error) {
	ip := c.GlobalString("ip")
	xbdm, err := goxbdm.NewXBDMClient(ip)
	if err != nil {
		_, err = xbdm.ChangeFreezeState(goxbdm.UnfrozenState)
	}

	return
}
