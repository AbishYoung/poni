/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * fileutil.go
 * Contains the file utility functions.
 *
 * @created 2023-08-20
 */
package main

import (
	"fmt"
	"os"
)

// fileExists
// Checks if a file exists.
// Returns true if it exists, false if not.
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// readFile
// Reads a file and returns the contents.
// Returns the contents of the file and an error.
func readFile(path string) ([]byte, error) {
	if !fileExists(path) {
		return nil, os.ErrNotExist
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file")
		}
	}(file)
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := stat.Size()
	data := make([]byte, size)
	_, err = file.Read(data)

	return data, err
}

// readStdin
// Reads from stdin and returns the contents.
// Returns the contents of stdin and an error.
func readStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	size := stat.Size()
	data := make([]byte, size)
	_, err = os.Stdin.Read(data)

	return data, err
}
