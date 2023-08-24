/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * ecdhke.go
 * Contains the functions for the ecdhke command.
 *
 * @created 2023-08-20
 */
package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"flag"
	"fmt"
)

func ECDHKEHandler(args []string) {
	var generateKeyPairs bool
	var generateSharedSecret bool
	var publicKey string
	var privateKey string
	var help bool

	command := flag.NewFlagSet("keyexchange", flag.ExitOnError)
	command.BoolVar(&generateKeyPairs, "generate-keys", false, "Generate initial key pairs")
	command.BoolVar(&generateSharedSecret, "generate-shared", false, "Generate shared secret")
	command.StringVar(&privateKey, "k", "", "Private key to use for key exchange")
	command.StringVar(&publicKey, "r", "", "Public key to use for key exchange")
	command.BoolVar(&help, "h", false, "Show help")
	err := command.Parse(args)
	if err != nil {
		fmt.Println("Error parsing arguments")
		return
	}

	if help {
		ECDHKEHelp()
		return
	}

	if generateKeyPairs {
		ECDHKEGenerateKeyPairs()
		return
	} else if generateSharedSecret {
		if publicKey == "" {
			ECDHKEHelp()
			return
		}

		if privateKey == "" {
			ECDHKEHelp()
			return
		}

		ECDHKEGenerateSharedSecret(privateKey, publicKey)
		return
	} else {
		ECDHKEHelp()
		return
	}
}

// ECDHKEHelp
// Prints the help message for the ecdhke command.
// Returns nothing.
func ECDHKEHelp() {
	fmt.Println("Usage: poni keyexchange [-generate-keys] [-generate-shared] [-k <private-key>] [-r <public-key>]")
	fmt.Println()
	fmt.Print(`        [-generate-keys]              Generate initial key pairs
        [-generate-shared]            Generate shared secret
                -k <key>              The private key to use for key exchange (required)
                -r <key>              The public key to use for key exchange (required)`)
	fmt.Println()
}

// ECDHKEGenerateKeyPairs
// Generates a private key and a public key.
// Returns nothing.
func ECDHKEGenerateKeyPairs() {
	curve := ecdh.X25519()
	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Failed to generate private key")
		return
	}

	publicKey := privateKey.PublicKey()
	fmt.Printf("Private key: %s\n", armor(privateKey.Bytes()))
	fmt.Printf("Public key: %s\n", armor(publicKey.Bytes()))
}

// ECDHKEGenerateSharedSecret
// Generates a shared secret from a private key and a public key.
// Returns nothing.
func ECDHKEGenerateSharedSecret(privateKey string, remoteKey string) {
	pKey, err := dearmor(privateKey)
	if err != nil {
		fmt.Println("Failed to decode private key")
		return
	}

	rKey, err := dearmor(remoteKey)
	if err != nil {
		fmt.Println("Failed to decode public key")
		return
	}

	curve := ecdh.X25519()
	pKeyStruct, err := curve.NewPrivateKey(pKey)
	if err != nil {
		fmt.Println("Failed to import private key")
		return
	}

	rKeyStruct, err := curve.NewPublicKey(rKey)
	if err != nil {
		fmt.Println("Failed to import public key")
		return
	}

	sharedSecret, _ := pKeyStruct.ECDH(rKeyStruct)
	if err != nil {
		fmt.Println("Failed to generate shared secret")
		return
	}

	fmt.Printf("Shared secret: %s\n", armor(sharedSecret))
}
