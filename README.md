# poni - simple file/stream encryption

---

Poni (stylized poni) is a file/stream encryption software meant to aid in the symmetric encryption of files, messages, and anything that can come through the standard input. This software is largely untested and no warranty is made about it's security. Using this software is to be done at your own risk. This is not to say that the code is unsafe or that the primitives that it relies on are insecure but rather that this specific repository has not been audited by anyone with the qualifications to do so.

Poni aims to protect your personal information using strong symmetric encryption in the form of AES-256 in GCM mode providing both message encryption and authentication in one go. This means that your ciphertext cannot be tampered with without failing the authentication check. Data integrity is important which is why it is not enough to encrypt alone. To this end poni bakes in authenticated encryption.

To make things easier on the end user poni includes a keygen which can both create a new key from a password or a passphrase or rederive a key from a known passphrase and salt. The keygen function leverages the Password Based Key Derivation Function 2 (PBKDF2) derivation function using the SHA3-256 digest algorithm at 200,000 iterations to derive a strong, secure key from a password or passphrase. Poni enforces a minimum password/passphrase length of 10 characters to provide the strongest keys possible.

## Usage

---
Poni currently supports the following commands, though more are planned for the future:
* encrypt
* decrypt
* keygen

Each of these commands have their own respective usages which can be accessed through the `-h` flag for each.

### Encrypt
This command provides message, file, and stdin encryption using a pre-generated key. The key can be generated using the keygen command. 

```
poni encrypt -k <key> [-p <message>] [-f <file>] [<stdin>]

-k <key>              The key to use for encryption
[-a]                  Armor the output using Base64 encoding
[-m <message>]        A message provided in plain text to encrypt
[-f <file>]           A file name whose contents will be encrypted
<nothing>             By default, poni will encrypt whatever comes
                      through stdin
```

### Decrypt
This command provides message, file, and stdin decryption using a pre-generated key. The key can be generated using the keygen command. 

```
poni decrypt -k <key> [-p <message>] [-f <file>] [<stdin>]

-k <key>              The key to use for decryption
[-s]                  Encode the output using string encoding
[-m <message>]        A message provided in plain text to decrypt
[-f <file>]           A file name whose contents will be derypted
<nothing>             By default, poni will decrypt whatever comes
                      through stdin
```

### Keygen
This command provides key generation and rederivation using a password or passphrase. The keygen command uses the PBKDF2 key derivation function with SHA3-256 as the digest algorithm and 200,000 iterations to derive a strong, secure key from a password or passphrase. The keygen command enforces a minimum password/passphrase length of 10 characters to provide the strongest keys possible.

```
poni keygen -p <password> [-s <salt>]

-p <password>         A password or passphrase to use for key generation
[-s <salt>]           A salt to use for key generation
```

## License
This software is licensed under the MIT license. See the LICENSE.md file for more information.