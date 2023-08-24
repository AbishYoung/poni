/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * pbkdf2.go
 * Contains the PBKDF2 key derivation function.
 *
 * @created 2023-08-20
 */
package main

import (
	"crypto/rand"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

// DeriveKey
// Derives a key from a password using PBKDF2.
// Returns the key, the salt, and an error.
func DeriveKey(password string) (key []byte, salt []byte, err error) {
	salt = make([]byte, 32)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	iterations := 200_000
	keySize := 32
	algo := sha3.New256
	key = pbkdf2.Key([]byte(password), salt, iterations, keySize, algo)
	return key, salt, nil
}

// DeriveKeyWithSalt
// Derives a key from a password and a salt using PBKDF2.
// Returns the key and an error.
func DeriveKeyWithSalt(password string, salt []byte) ([]byte, error) {
	iterations := 200_000
	keySize := 32
	algo := sha3.New256
	key := pbkdf2.Key([]byte(password), salt, iterations, keySize, algo)
	return key, nil
}
