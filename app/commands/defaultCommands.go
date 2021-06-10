package commands

import "github.com/urfave/cli"

type DefaultCommandsInterf interface {
	AddCommands(commands ...cli.Command)
}

type DefaultCommands struct {
	CommandList []cli.Command
}

func (d *DefaultCommands) AddCommands(commands ...cli.Command) {
	d.CommandList = append(d.CommandList, commands...)
}
