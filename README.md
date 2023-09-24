# CipherCloak
This program encrypts and decrypts plaintext using a system similar to the [one-time pad](https://en.wikipedia.org/wiki/One-time_pad).  
## Technologies
* Go
## GIF Walkthrough
<p>
<image src = "ciphercloak_gif.gif" title = "CipherCloak GIF"><br>
</p>

## How to Run
* Clone this repository.
* `cd` into the directory.
* Run `make` to compile the project. 
## Project Details
* Overview
    * Plaintext is the information you want to encrypt.  Allowed characters include the capital letters of the English alphabet and a space.
    * Key is a randomly generated sequence of characters that are used to encrypt plaintext and to decrypt ciphertext. Like plaintext, allowed characters include the capital letters of the English alphabet and a space.
    * Ciphertext is the encrypted version of plaintext using the key.
* Keymaker
    * The keymaker generates a key whose length is determined by the argument passed to it. The following command will produce a key of length 50 and store it to a file named 'key': `./keymaker 50 > key`.
* Encrypt Server
    * The encrypt server receives the names of plaintext and key from the encrypt client, opens the files, encrypts the plaintext using a system similar to [one-time pad](https://en.wikipedia.org/wiki/One-time_pad), and sends the ciphertext back to the client. The following command will start the encrypt server on port 12345: `./encrypt_server 12345 &`.
* Encrypt Client
    * The encrypt client connects to the encrypt server, sends the names of plaintext and key, and asks it to perform encryption on it. Once the encrypt server is done encrypting the plaintext, encrypt client receives the ciphertext from the server and outputs it to stdout. The following command will send the plaintext and key to the encrypt server (open on port 12345) and output the encrypted version to a file named 'ciphertext': `./encrypt_client plaintext key 12345 > ciphertext`.
* Decrypt Server
    * The decrypt server receives the names of cipher and key from the decrypt client, opens the files, decrypts the cipher using the key, and sends the decrypted text (plaintext) back to the client. The following command will start the decrypt server on port 54321: `./decrypt_server 54321 &`.
* Decrypt Client
    * The decrypt client connects to the decrypt server, sends the names of cipher and key, and asks it to perform decryption on it. Once the decrypt server is done decrypting the cipher, decrypt client receives the decrypted text (plaintext) from the server and outputs it to stdout. The following command will send the cipher and key to the decrypt server (open on port 54321) and output the decrypted version to a file named 'plaintext_decrypted': `./decrypt_client plaintext key 54321 > plaintext_decrypted`.
## How to Test using test.sh
* `cd` into the project directory.
* Run `chmod +x test.sh` to make the test script executable.
* Run `make` to complie the project.
* Run `./test.sh` to run the test script.
* If you want to run a clean-up script, run `chmod +x clear.sh` followed by `./clear.sh`.
## Contact
Elliott Larsen
* Email elliottlrsn@gmail.com
* GitHub [@elliottlarsen](https://github.com/elliottlarsen)
* LinkedIn [@elliottlarsen](https://www.linkedin.com/in/elliottlarsen)
