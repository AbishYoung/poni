/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * keygen.go
 * Contains the keygen command.
 *
 * @created 2023-08-23
 */
package main

import (
	"flag"
	"fmt"
)

// keygenHandler
// Handles the keygen command.
// Returns nothing.
func keygenHandler(args []string) {
	var password string
	var salt string
	var help bool

	command := flag.NewFlagSet("keygen", flag.ExitOnError)
	command.StringVar(&password, "p", "", "The password to use for key generation")
	command.StringVar(&salt, "s", "", "The salt to use for key generation")
	command.BoolVar(&help, "h", false, "Show the keygen help message")
	err := command.Parse(args)
	if err != nil {
		return
	}

	if help {
		keygenUsage()
		return
	}

	if password == "" {
		keygenUsage()
		return
	}

	if salt == "" {
		if len(password) < 10 {
			fmt.Println("For your security passwords must be at least 10 characters long")
			return
		}

		key, keySalt, err := deriveKey(password)
		if err != nil {
			fmt.Println("Failed to generate key")
			return
		}

		fmt.Println("Key:", armor(key))
		fmt.Println("Salt:", armor(keySalt))
	} else {
		decodedSalt, err := dearmor(salt)
		if err != nil {
			fmt.Println("Failed to decode salt")
			return
		}

		key, err := deriveKeyWithSalt(password, decodedSalt)
		if err != nil {
			fmt.Println("Failed to generate key")
			return
		}

		fmt.Println("Key:", armor(key))
	}

}

// keygenUsage
// Prints the keygen help message.
// Returns nothing.
func keygenUsage() {
	fmt.Println("Usage: poni keygen -p <password> [-s <salt>]")
	fmt.Println()
	fmt.Println(`        -p <password>           A password or passphrase to use for key generation
        [-s <salt>]           A salt to use for key generation`)
	fmt.Println()
}
