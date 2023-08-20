# poni
Poni is a simple encryption utility meant to make the encryption and authentication of small files easier (since OpenSSL is hesitant about adding AEAD ciphers). This software is mostly untested, unverified, and should not be understood to be bulletproof or abundantly secure. With that said, I mean it does encrypt and the code is open to anyone so if there is someone out there who can commit some fixes or improvements that would be great.

Poni secures data using a password which is passed through 200,000 iterations of the PBKDF2 derivation function using the SHA3 256 hash. The password as well as the data are given a unique salt and nonce respectively. Password reuse should be kept to a minimum though. The data is then passed through AES-256 in GCM mode which is in turn used to authenticate the integrity of the ciphertext upon decryption ensuring that it has not been tampered with.
## usage
#### encrypt
`poni encrypt -p PASSWORD -i INPUT_FILE -o OUTPUT_FILE`
#### decrypt
`poni decrypt -p PASSWORD -i INPUT_FILE -o OUTPUT_FILE`
#### generate a key
`poni derive-key -p PASSWORD`