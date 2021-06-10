package commands

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GenerateCommands(defaultCommands *DefaultCommands) {
	flags := generateFlags()
	ipAction, serverAction := generateActions()

	ip := cli.Command{
		Name:   "ip",
		Usage:  "Search an specific ip given an web address",
		Flags:  flags,
		Action: ipAction,
	}

	server := cli.Command{
		Name:   "server",
		Usage:  "Find internet servers",
		Flags:  flags,
		Action: serverAction,
	}

	defaultCommands.AddCommands(ip, server)
}

func generateFlags() (flags []cli.Flag) {
	host := cli.StringFlag{
		Name:  "host",
		Value: "amazon.com.br",
	}

	flags = []cli.Flag{
		host,
	}
	return
}

func generateActions() (ipAction interface{}, serverAction interface{}) {
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

	serverAction = func(c *cli.Context) {
		host := c.String("host")

		servers, err := net.LookupNS(host)

		if err != nil {
			log.Fatal(err)
		}

		for _, server := range servers {
			fmt.Println(server.Host)
		}
	}

	return
}
