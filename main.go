package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rara-ch/myknowledge.git/command"
)

func main() {
	cmds := command.NewCLICommands()

	cmds.Register("hello", helloHandler)
	cmds.Register("goodbye", goodbyeHandler)
	cmds.Register("error", errorHandler)

	input := os.Args

	if len(input) < 2 {
		fmt.Println("Please enter a command for myknowledge to work")
		return
	}

	cmd := command.InputCommand{
		Name: input[1],
		Args: input[2:],
	}

	err := cmds.Run(cmd)
	if err != nil {
		fmt.Println(err)
	}
}

func helloHandler(cmd command.InputCommand) error {
	fmt.Println("Hello World")
	return nil
}

func goodbyeHandler(cmd command.InputCommand) error {
	fmt.Println("Goodbye Cruel World")
	return nil
}

func errorHandler(cmd command.InputCommand) error {
	fmt.Println("error handler")
	return errors.New("this is an error")
}
