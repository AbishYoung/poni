# poni - simple file/stream encryption


Poni (stylized poni) is a file/stream encryption software meant to aid in the symmetric encryption of files, messages, and anything that can come through the standard input. This software is largely untested and no warranty is made about its security. Using this software is to be done at your own risk. This is not to say that the code is unsafe or that the primitives that it relies on are insecure but rather that this specific repository has not been audited by anyone with the qualifications to do so.

Poni aims to protect your personal information using strong symmetric encryption in the form of AES-256 in GCM mode providing both message encryption and authentication in one go. This means that your ciphertext cannot be tampered with without failing the authentication check. Data integrity is important which is why it is not enough to encrypt alone. To this end poni bakes in authenticated encryption.

To make things easier on the end user poni includes a keygen which can both create a new key from a password or a passphrase or rederive a key from a known passphrase and salt. The keygen function leverages the scrypt key derivation function with an N-factor of 131,072 to derive a strong, secure key from a password or passphrase. Poni enforces a minimum password/passphrase length of 10 characters to provide the strongest keys possible.

## Usage

Poni currently supports the following commands, though more are planned for the future:
* encrypt
* decrypt
* keygen
* keyexchange

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
This command provides key generation and rederivation using a password or passphrase. The keygen command uses the scrypt key derivation function with an N-factor of 131,072 to derive a strong, secure key from a password or passphrase. The keygen command enforces a minimum password/passphrase length of 10 characters to provide the strongest keys possible.

```
poni keygen -p <password> [-s <salt>]

-p <password>         A password or passphrase to use for key generation
[-s <salt>]           A salt to use for key generation
```

### Keyexchange
This command provides a way to generate a shared secret through an unsecured channel using the Elliptical Curve Diffie-Hellman Key Exchange (ECDHKE) algorithm on the X25519 curve. This command is meant to be used in conjunction with the encrypt and decrypt commands to provide a way to encrypt and decrypt messages without having to share a key beforehand. The keyexchange command will output a public key which can be shared with the other party. The other party will then provide their public key to the keyexchange command which will then output a shared secret which can be used as a key for the encrypt and decrypt commands.

```
Usage: poni keyexchange [-generate-keys] [-generate-shared] [-k <private-key>] [-r <public-key>]

[-generate-keys]              Generate initial key pairs
[-generate-shared]            Generate shared secret
        -k <key>              The private key to use for key exchange (required)
        -r <key>              The public key to use for key exchange (required)
```

A complete key exchange would look something like this:

```
# Alice
$ poni keyexchange -generate-keys
Private key: XjhGxThmlIoQk27qgWnuAqzHEWxFk9U0SWgtPlCZLIE=
Public key: Ca5tFXzqUXkGGBcPzT6fNBZfTwirKPtSNU0BA+uKz38=

# Bob
$ poni keyexchange -generate-keys
Private key: gmZsULnQPPPIVa8oClCd9XvV8v3JM/HU94YNfW8z1jc=
Public key: JuOQ5C/s3SSawsT1jTzYW8s1hk9ahaB2iGiHu3RMWG8=

# Alice
$ poni keyexchange -k XjhGxThmlIoQk27qgWnuAqzHEWxFk9U0SWgtPlCZLIE= -r JuOQ5C/s3SSawsT1jTzYW8s1hk9ahaB2iGiHu3RMWG8=
Shared secret: m5YRm+cz5lKdSfQ6gdzsv8MajXBO/LYqD0lhZKrcJQA=

# Bob
$ poni keyexchange -k gmZsULnQPPPIVa8oClCd9XvV8v3JM/HU94YNfW8z1jc= -r Ca5tFXzqUXkGGBcPzT6fNBZfTwirKPtSNU0BA+uKz38=
Shared secret: m5YRm+cz5lKdSfQ6gdzsv8MajXBO/LYqD0lhZKrcJQA=
```

## License
This software is licensed under the MIT license. See the LICENSE.md file for more information.