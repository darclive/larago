package cli

import (
	"os"

	"github.com/lara-go/larago"
	"github.com/lara-go/larago/foundation/bootstrappers"
	"github.com/urfave/cli"
)

// Kernel for cli commands.
type Kernel struct {
	Application *larago.Application

	bootstrappers []larago.Bootstrapper
	commands      []larago.ConsoleCommand
}

// NewKernel constructor.
func NewKernel() *Kernel {
	kernel := &Kernel{}

	kernel.SetBootstrappers(
		bootstrappers.DetectEnv,
		bootstrappers.LoadConfig,
		bootstrappers.BootProviders,
	)

	return kernel
}

// SetBootstrappers sets Application bootstrappers.
func (k *Kernel) SetBootstrappers(bootstrappers ...larago.Bootstrapper) {
	k.bootstrappers = bootstrappers
}

// Handle console commands.
func (k *Kernel) Handle() {
	if err := k.Application.BootstrapWith(k.bootstrappers...); err != nil {
		panic(err)
	}

	app := cli.NewApp()

	app.Version = k.Application.Version
	app.Name = k.Application.Name
	app.Usage = k.Application.Description

	app.Flags = k.getGlobalFlags()
	app.Commands = k.makeCommands(k.Application.GetCommands())

	app.Run(os.Args)
}

// GetGlobalFlags registers global flags.
func (k *Kernel) getGlobalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "home, H",
			Value:       k.Application.HomeDirectory,
			Usage:       "path to home directory",
			Destination: &k.Application.HomeDirectory,
		},
	}
}

// Cli commands factory.
func (k *Kernel) makeCommands(commands []larago.ConsoleCommand) []cli.Command {
	var cliCommands []cli.Command

	for _, command := range commands {
		cliCommands = append(cliCommands, k.makeCommand(command))
	}

	return cliCommands
}

// Make command for the cli package.
func (k *Kernel) makeCommand(command larago.ConsoleCommand) cli.Command {
	cliCommand := command.GetCommand()

	// Cli command handler.
	cliCommand.Action = func(c *cli.Context) error {
		// Resolve command's dependencies.
		k.Application.Make(command)

		// Run Handler.
		if err := command.Handle(c.Args()); err != nil {
			panic(err)
		}

		return nil
	}

	return cliCommand
}
