/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * encrypt.go
 * Contains the encryption command functions.
 *
 * @created 2023-08-23
 */
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
)

// encryptHandler
// Handles the encrypt command.
// Returns nothing.
func encryptHandler(args []string) {
	var mode string
	var key string
	var plaintext string
	var file string
	var help bool
	var armorOutput bool
	var ciphertext []byte
	var iv []byte

	command := flag.NewFlagSet("encrypt", flag.ExitOnError)
	command.StringVar(&key, "k", "", "The key to use for encryption")
	command.StringVar(&plaintext, "m", "", "The plaintext to encrypt")
	command.StringVar(&file, "f", "", "The file to encrypt")
	command.BoolVar(&armorOutput, "a", false, "Whether to armor the output")
	command.BoolVar(&help, "h", false, "Show the encryption help message")
	err := command.Parse(args)
	if err != nil {
		return
	}

	if help {
		encryptUsage()
		return
	}

	// change mode based on present flags
	if plaintext != "" {
		mode = "plaintext"
	} else if file != "" {
		mode = "file"
	} else {
		mode = "stdin"
	}

	// check if key is present
	if key == "" {
		encryptUsage()
		return
	}

	decodedKey, err := dearmor(key)
	if err != nil {
		fmt.Println("Failed to decode key")
		return
	}

	switch mode {
	case "plaintext":
		ciphertext, iv, err = encrypt(decodedKey, []byte(plaintext))
	case "file":
		plaintext, err := readFile(file)
		if err != nil {
			fmt.Println("Failed to read file")
			return
		}
		ciphertext, iv, err = encrypt(decodedKey, plaintext)
	case "stdin":
		plaintext, err := readStdin()
		if err != nil {
			fmt.Println("Failed to read stdin")
			return
		}
		ciphertext, iv, err = encrypt(decodedKey, plaintext)
	}

	if err != nil {
		fmt.Println("Failed to encrypt")
		return
	}

	if armorOutput {
		fmt.Println(armor(append(iv, ciphertext...)))
	} else {
		fmt.Println(append(iv, ciphertext...))
	}
}

func encryptUsage() {
	fmt.Println("Usage: poni encrypt -k <key> [-a armor] [-p <plaintext>] [-f <file>] [stdin]")
	fmt.Println()
	fmt.Print(`        -k <key>              The key to use for encryption
        [-a]                  Armor the output using Base64 encoding
        [-m <message>]        A message provided in plain text to encrypt
        [-f <file>]           A file name whose contents will be encrypted
        <nothing>             By default, poni will encrypt whatever comes
                              through stdin`)
	fmt.Println()
}

// encrypt
// Encrypts a plaintext using AES-GCM.
// Returns the ciphertext, the IV, and an error.
func encrypt(key []byte, plaintext []byte) (ciphertext []byte, iv []byte, err error) {
	iv = make([]byte, 12)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	AESCipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	ciphertext = AESCipher.Seal(nil, iv, plaintext, nil)

	return ciphertext, iv, nil
}
