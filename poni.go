package main

import (
	"fmt"
	"os"
)

// main
// The main function.
// Returns nothing.
func main() {
	registerCommand("encrypt", encryptHandler)
	registerCommand("decrypt", decryptHandler)
	registerCommand("keygen", keygenHandler)

	if len(os.Args) < 2 {
		help()
		return
	}

	args := os.Args[1:]
	if err := handleCommand(args[0], args[1:]); err != nil {
		fmt.Println(err)
	}
}

// help
// Prints the help message.
// Returns nothing.
func help() {
	fmt.Println("Usage: poni [command] [args]")
	fmt.Println("Commands:")
	for _, command := range Commands {
		fmt.Printf("  %s\n", command.Name)
	}
	fmt.Println("Use poni [command] -h for more information about a command")
}
