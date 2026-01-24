package main

import (
	"fmt"
)

type inputCommand struct {
	name string
	args []string
}

type handler func(*state, inputCommand) error

type cliCommands struct {
	handlers map[string]handler
}

func newCLICommands() cliCommands {
	return cliCommands{
		handlers: make(map[string]handler),
	}
}

func (c *cliCommands) run(state *state, inputCommand inputCommand) error {
	if handler, ok := c.handlers[inputCommand.name]; !ok {
		return fmt.Errorf("%s is not not a valid myknowledge command", inputCommand.name)
	} else {
		return handler(state, inputCommand)
	}
}

func (c *cliCommands) register(name string, handler handler) {
	c.handlers[name] = handler
}
