package app

import (
	c "address-searcher/app/commands"
	"github.com/urfave/cli"
)

// Generate Return the CLI application ready to be executed
func Generate() *cli.App {
	defaultCommands := c.DefaultCommands{}
	c.GenerateCommands(&defaultCommands)

	app := cli.NewApp()
	app.Name = "Address Searcher"
	app.Usage = "Search Ip and Logs"
	app.Commands = defaultCommands.CommandList
	return app
}

