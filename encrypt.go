/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * encrypt.go
 * Contains the encryption and decryption functions.
 *
 * @created 2023-08-20
 */
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// Encrypt
// Encrypts a plaintext using AES-GCM.
// Returns the ciphertext, the IV, and an error.
func Encrypt(key []byte, plaintext []byte) (ciphertext []byte, iv []byte, err error) {
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

// Decrypt
// Decrypts a ciphertext using AES-GCM.
// Returns the plaintext and an error.
func Decrypt(key []byte, ciphertext []byte, iv []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	AESCipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = AESCipher.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
