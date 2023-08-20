package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"os"
)

func main() {
	/**
	 * Command-line interface
	 *
	 * Actions:
	 *     derive-key -p [PASSWORD]
	 *     encrypt:
	 *  	   -p [PASSWORD] -i [INPUT] -o [OUTPUT] // uses default key derivation parameters
	 *     decrypt:
	 *  	   -p [PASSWORD] -i [INPUT] -o [OUTPUT] // uses salt and iterations from encoded input file
	 */
	var password string
	var input string
	var output string

	// parse command-line arguments
	deriveKeyCmd := flag.NewFlagSet("derive-key", flag.ExitOnError)
	deriveKeyCmd.StringVar(&password, "p", "", "password")

	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptCmd.StringVar(&password, "p", "", "password")
	encryptCmd.StringVar(&input, "i", "", "input file")
	encryptCmd.StringVar(&output, "o", "", "output file")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptCmd.StringVar(&password, "p", "", "password")
	decryptCmd.StringVar(&input, "i", "", "input file")
	decryptCmd.StringVar(&output, "o", "", "output file")

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "derive-key":
		err := deriveKeyCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		key, salt := deriveKey(password)
		fmt.Println("Key: ", base64.StdEncoding.EncodeToString(key))
		fmt.Println("Salt: ", base64.StdEncoding.EncodeToString(salt))
	case "encrypt":
		err := encryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		encrypt(password, input, output)
	case "decrypt":
		err := decryptCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		decrypt(password, input, output)
	default:
		help()
		os.Exit(1)
	}
}

func help() {
	fmt.Println("Usage: poni [COMMAND] [OPTIONS]")
	fmt.Println("Commands:")
	fmt.Println("    derive-key -p [PASSWORD]")
	fmt.Println("    encrypt -p [PASSWORD] -i [INPUT] -o [OUTPUT]")
	fmt.Println("    decrypt -p [PASSWORD] -i [INPUT] -o [OUTPUT]")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func deriveKeyWithSalt(password string, salt []byte) ([]byte, []byte) {
	iterations := 200_000
	keySize := 32
	algo := sha3.New256
	key := pbkdf2.Key([]byte(password), salt, iterations, keySize, algo)
	return key, salt
}

func deriveKey(password string) ([]byte, []byte) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return deriveKeyWithSalt(password, salt)
}

// @TODO: not endianess safe
func readCipherFile(filename string) ([]byte, []byte, []byte) {
	if !fileExists(filename) {
		fmt.Println("Input file does not exist.")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileSize := fileStat.Size()

	// read salt
	salt := make([]byte, 32)
	_, err = file.Read(salt)

	// read nonce
	nonce := make([]byte, 12)
	_, err = file.Read(nonce)

	// read ciphertext
	ciphertext := make([]byte, fileSize-44)
	_, err = file.Read(ciphertext)

	return salt, nonce, ciphertext
}

func encrypt(password string, input string, output string) {
	if !fileExists(input) {
		fmt.Println("Input file does not exist.")
		os.Exit(1)
	}
	plaintext, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	key, salt := deriveKey(password)
	nonce := make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	AESGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ciphertext := AESGCM.Seal(nil, nonce, plaintext, nil)

	// write salt, nonce, and ciphertext to file
	var buf bytes.Buffer

	buf.Write(salt)
	buf.Write(nonce)
	buf.Write(ciphertext)

	err = ioutil.WriteFile(output, buf.Bytes(), 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func decrypt(password string, input string, output string) {
	salt, nonce, ciphertext := readCipherFile(input)
	key, _ := deriveKeyWithSalt(password, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	AESGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := AESGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(output, plaintext, 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
