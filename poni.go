package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var password string
	var key string
	var salt string
	var encodeOutput bool
	var stringOutput bool

	// keygen flags
	keygenCmd := flag.NewFlagSet("keygen", flag.ExitOnError)
	keygenCmd.StringVar(&password, "p", "", "password")
	keygenCmd.StringVar(&salt, "s", "", "salt")

	// encrypt flags
	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptCmd.StringVar(&key, "k", "", "encryption key")
	encryptCmd.BoolVar(&encodeOutput, "e", false, "base64 encode output")

	// decrypt flags
	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptCmd.StringVar(&key, "k", "", "encryption key")
	decryptCmd.BoolVar(&encodeOutput, "e", false, "base64 encode output")
	decryptCmd.BoolVar(&stringOutput, "s", false, "output as string")

	// require action
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("action is required: keygen, encrypt, decrypt")
		return
	}

	switch os.Args[1] {
	case "keygen":
		err := keygenCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}

		if password == "" {
			fmt.Println("password is required")
			flag.Usage()
			return
		}

		if salt != "" {
			decodedSalt, err := b64d(salt)
			if err != nil {
				fmt.Println(err)
				return
			}
			derivedKey, err := DeriveKeyWithSalt(password, decodedSalt)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("key: %s\n", b64e(derivedKey))
			fmt.Printf("salt: %s\n", b64e(decodedSalt))
			return
		} else {
			derivedKey, salt, err := DeriveKey(password)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("key: %s\n", b64e(derivedKey))
			fmt.Printf("salt: %s\n", b64e(salt))
		}
	case "encrypt":
		err := encryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}

		if key == "" {
			fmt.Println("key is required")
			flag.Usage()
			return
		}

		decodedKey, err := b64d(key)
		if err != nil {
			fmt.Println(err)
			return
		}

		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}

		plaintext := stdin
		ciphertext, iv, err := Encrypt(decodedKey, plaintext)
		if err != nil {
			fmt.Println(err)
			return
		}

		blob := WriteCiphertext(iv, ciphertext)

		if encodeOutput {
			fmt.Println(b64e(blob))
		} else {
			fmt.Println(blob)
		}
	case "decrypt":
		err := decryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}

		if key == "" {
			fmt.Println("key is required")
			flag.Usage()
			return
		}

		decodedKey, err := b64d(key)
		if err != nil {
			fmt.Println(err)
			return
		}

		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}

		ciphertext := stdin
		iv, ciphertext, err := ReadCiphertext(ciphertext)
		if err != nil {
			fmt.Println(err)
			return
		}

		plaintext, err := Decrypt(decodedKey, ciphertext, iv)
		if err != nil {
			fmt.Println(err)
			return
		}

		if encodeOutput {
			fmt.Println(b64e(plaintext))
		} else if stringOutput {
			fmt.Println(string(plaintext))
		} else {
			fmt.Println(plaintext)
		}
	default:
		flag.Usage()
		return
	}
}
