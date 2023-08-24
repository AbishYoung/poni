/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * decrypt.go
 * Contains the decryption command functions.
 *
 * @created 2023-08-20
 */
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"flag"
	"fmt"
)

// decryptHandler
// Handles the decrypt command.
// Returns nothing.
func decryptHandler(args []string) {
	var mode string
	var key string
	var ciphertext string
	var file string
	var help bool
	var asString bool
	var plaintext []byte
	var iv []byte

	command := flag.NewFlagSet("decrypt", flag.ExitOnError)
	command.StringVar(&key, "k", "", "The key to use for decryption")
	command.StringVar(&ciphertext, "m", "", "The plaintext to decrypt")
	command.StringVar(&file, "f", "", "The file to decrypt")
	command.BoolVar(&asString, "s", false, "Whether to output the plaintext as a string")
	command.BoolVar(&help, "h", false, "Show the decryption help message")
	err := command.Parse(args)
	if err != nil {
		return
	}

	if help {
		decryptUsage()
		return
	}

	// change mode based on present flags
	if ciphertext != "" {
		mode = "plaintext"
	} else if file != "" {
		mode = "file"
	} else {
		mode = "stdin"
	}

	// check if key is present
	if key == "" {
		decryptUsage()
		return
	}

	decodedKey, err := dearmor(key)
	if err != nil {
		fmt.Println("Failed to decode key")
		return
	}

	switch mode {
	case "plaintext":
		blob, err := dearmor(ciphertext)
		if err != nil {
			fmt.Println("Failed to decode ciphertext")
			return
		}
		iv = blob[:12]
		ciphertext := blob[12:]
		plaintext, err = decrypt(decodedKey, ciphertext, iv)
	case "file":
		blob, err := readFile(file)
		if err != nil {
			fmt.Println("Failed to read file")
			return
		}
		iv = blob[:12]
		ciphertext := blob[12:]
		plaintext, err = decrypt(decodedKey, ciphertext, iv)
	case "stdin":
		blob, err := readStdin()
		if err != nil {
			fmt.Println("Failed to read stdin")
			return
		}
		iv = blob[:12]
		ciphertext := blob[12:]
		plaintext, err = decrypt(decodedKey, ciphertext, iv)
	}

	if err != nil {
		fmt.Println("Failed to decrypt")
		return
	}

	if asString {
		fmt.Println(string(plaintext))
	} else {
		fmt.Println(plaintext)
	}
}

func decryptUsage() {
	fmt.Println("Usage: poni decrypt -k <key> [-a armor] [-p <plaintext>] [-f <file>] [stdin]")
	fmt.Println()
	fmt.Print(`        -k <key>              The key to use for decryption
        [-s]                  Encode the output using string encoding
        [-m <message>]        A message provided in plain text to decrypt
        [-f <file>]           A file name whose contents will be decrypted 
        <nothing>             By default, poni will decrypt whatever comes
                              through stdin`)
	fmt.Println()
}

// decrypt
// Decrypts a ciphertext using AES-GCM.
// Returns the plaintext and an error.
func decrypt(key []byte, ciphertext []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	AESCipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := AESCipher.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
