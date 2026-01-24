package command

import "fmt"

type InputCommand struct {
	Name string
	Args []string
}

type handler func(InputCommand) error

type cLICommands struct {
	handlers map[string]handler
}

func NewCLICommands() cLICommands {
	return cLICommands{
		handlers: make(map[string]handler),
	}
}

func (c *cLICommands) Run(inputCommand InputCommand) error {
	if handler, ok := c.handlers[inputCommand.Name]; !ok {
		return fmt.Errorf("%s is not not a valid myknowledge command", inputCommand.Name)
	} else {
		return handler(inputCommand)
	}
}

func (c *cLICommands) Register(name string, handler handler) {
	c.handlers[name] = handler
}
