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

import "encoding/base64"

// ReadCiphertext
// Reads a ciphertext blob and returns the IV and ciphertext.
// Returns the IV, the ciphertext, and an error.
func ReadCiphertext(blob []byte) (iv []byte, ciphertext []byte, err error) {
	blob, err = b64d(string(blob))
	if err != nil {
		return nil, nil, err
	}

	iv = blob[:12]
	ciphertext = blob[12:]
	return iv, ciphertext, nil
}

// WriteCiphertext
// Writes a ciphertext blob from an IV and ciphertext.
// Returns the ciphertext blob.
func WriteCiphertext(iv []byte, ciphertext []byte) (blob []byte) {
	blob = append(iv, ciphertext...)
	return blob
}

// b64e
// Encodes a byte slice to base64.
// Returns the encoded string.
func b64e(data []byte) (encoded string) {
	encoded = base64.StdEncoding.EncodeToString(data)
	return encoded
}

// b64d
// Decodes a base64 string to a byte slice.
// Returns the decoded byte slice and an error.
func b64d(data string) (decoded []byte, err error) {
	decoded, err = base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
