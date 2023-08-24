/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * yggdrasil.go
 * Contains the code that handles the command line arguments
 *
 * @created 2023-08-23
 */
package main

import (
	"errors"
	"fmt"
)

// Commands
// Contains all the commands.
var Commands []Command

// Command
// Represents a command and it's handler.
type Command struct {
	Name    string
	Handler func([]string)
}

func registerCommand(name string, handler func([]string)) {
	Commands = append(Commands, Command{Name: name, Handler: handler})
}

func findHandler(name string) (func([]string), error) {
	for _, command := range Commands {
		if command.Name == name {
			return command.Handler, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Command %s not found", name))
}

func handleCommand(name string, args []string) error {
	handler, err := findHandler(name)
	if err != nil {
		return err
	}

	handler(args)
	return nil
}
