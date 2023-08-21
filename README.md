# poni
Poni is a simple encryption utility meant to make the encryption and authentication of small files easier (since OpenSSL is hesitant about adding AEAD ciphers). This software is mostly untested, unverified, and should not be understood to be bulletproof or abundantly secure. With that said, I mean it does encrypt and the code is open to anyone so if there is someone out there who can commit some fixes or improvements that would be great.

Poni secures data using a password which is passed through 200,000 iterations of the PBKDF2 derivation function using the SHA3 256 hash. The password as well as the data are given a unique salt and nonce respectively. Password reuse should be kept to a minimum though. The data is then passed through AES-256 in GCM mode which is in turn used to authenticate the integrity of the ciphertext upon decryption ensuring that it has not been tampered with.
## Usage
### keygen
Before you can encrypt you need to generate a key. You can do this with the `keygen` action as so:

`poni keygen -p test`

Doing so will produce both a Base64 encoded key as well as a salt that can then be used in encryption/decryption or to rederive the same key. If you do not specify a salt then a new one will be randomly generated for you using a cryptographically secure PRNG.

To rederive the key you can use the `keygen` action again with the `-s` flag to specify the salt and the `-p` flag to specify the password. For example:

`poni keygen -s 5f2a2b1c -p test`

### encrypt
To encrypt a file you can use the `encrypt` action as so:

`echo "Hello, World!" | poni encrypt -k key`

Doing so will produce a Base64 encoded ciphertext as well as a Base64 encoded nonce that can then be used in decryption. You can control the way that the output is displayed using the `-b` flag to specify that the ciphertext block should be Base64 encoded. You can also pipe the output to a file or to another program using standard pipes.

### decrypt
To decrypt a file you can use the `decrypt` action as so:

`echo "ciphertext" | poni decrypt -k key`

Doing so will produce the original plaintext. You can control the way that the input is read using the `-b` flag to specify that the ciphertext block should be Base64 encoded or the `-s` flag to specify that the plaintext should be encoded as a UTF-8 string (this is useful for messages). You can also pipe the output to a file or to another program using standard pipes.

## License
This software is licensed under the MIT license. See the LICENSE.md file for more information.