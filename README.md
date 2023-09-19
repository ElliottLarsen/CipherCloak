# CipherCloak
This program encrypts and decrypts plaintext using a system similar to the one-time pad.  
## Technologies
* Go
## GIF Walkthrough

## How to Run
* Clone this repository.
* `cd` into the directory.
* Run `make` to compile the project. 
## Project Details
* Overview
    * Plaintext is the information you want to encrypt.  Allowed characters include the capital letters of the English alphabet and a space.
    * Key is a randomly generated sequence of characters that are used to encrypt plaintext and to decrypt ciphertext.  Like plaintext, allowed characters include the capital letters of the English alphabet and a space.
    * Ciphertext is the encrypted version of plaintext using the key.
* Keymaker
    * The keymaker generates a key whose length is determined by the argument passed to it. The following command will produce a key of length 50 and store it to a file named 'key': `./keymaker 50 > key`.
