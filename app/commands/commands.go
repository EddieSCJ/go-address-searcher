package commands

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

func GenerateCommands(defaultCommands *DefaultCommands) {
	ipFlags := generateFlags()
	ipAction := generateActions()

	ip := cli.Command{
		Name:   "ip",
		Usage:  "Search an specific ip given an web address",
		Flags:  ipFlags,
		Action: ipAction,
	}

	defaultCommands.AddCommands(ip)
}

func generateFlags() (ipFlags []cli.Flag) {
	host := cli.StringFlag{
		Name:  "host",
		Value: "amazon.com.br",
	}

	ipFlags = []cli.Flag{
		host,
	}
	return
}

func generateActions() (ipAction interface{}) {
	ipAction = func(c *cli.Context) {
		host := c.String("host")

		ips, err := net.LookupIP(host)
		if err != nil {
			log.Fatal(err)
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

	return
}
