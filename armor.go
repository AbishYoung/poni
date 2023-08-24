/**
 * This file is a part of the poni project and is licensed under the MIT license.
 * See LICENSE.md for details.
 *
 * armor.go
 * Contains the armor and dearmor functions.
 *
 * @created 2023-08-23
 */
package main

import "encoding/base64"

// armor
// Encodes a byte slice to base64.
// Returns the encoded string.
func armor(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// dearmor
// Decodes a base64 string to a byte slice.
// Returns the decoded byte slice and an error.
func dearmor(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
