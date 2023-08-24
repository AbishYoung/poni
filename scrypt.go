/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * scrypt.go
 * Contains the scrypt key derivation function.
 *
 * @created 2023-08-23
 */
package main

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

// deriveKey
// Derives a key from a password using scrypt.
// Returns the key, the salt, and an error.
func deriveKey(password string) (key []byte, salt []byte, err error) {
	salt = make([]byte, 32)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	key, err = scrypt.Key([]byte(password), salt, 131072, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
}

// deriveKeyWithSalt
// Derives a key from a password and a salt using scrypt.
// Returns the key and an error.
func deriveKeyWithSalt(password string, salt []byte) ([]byte, error) {
	key, err := scrypt.Key([]byte(password), salt, 131072, 8, 1, 32)
	if err != nil {
		return nil, err
	}

	return key, nil
}
