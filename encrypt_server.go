package main

import (
  "fmt"
  "io/ioutil"
  "net"
  "os"
)

func main() {
  // Check for the command line argument.
  if len(os.Args) != 2 {
    fmt.Fprintln(os.Stderr, "Incorrect command line arguments.  It should be ./server <port>")
    os.Exit(0)
  }

  port := os.Args[1]
  // Create a TCP listener on the port.
  listener, err := net.Listen("tcp", ":" + port)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error listening on port " + port)
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  defer listener.Close()

  fmt.Fprintln(os.Stdin, "Server is listening on port " + port)

  for {
    // Accept a client connection.
    conn, err := listener.Accept()
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error accepting connection.")
      fmt.Fprintln(os.Stderr, err)
      continue
    }
    // Handle the client connection in a goroutine.
    go encrypt_text(conn)
  }
}

func encrypt_text(conn net.Conn) {
  defer conn.Close()

  // Read file names from the client.
  buffer := make([]byte, 200000)
  n, err := conn.Read(buffer)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading file names from the client.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  file_names := string(buffer[:n])
  file_names_arr := split_file_names(file_names)

  // Open plaintext and key files.
  plaintext, err := ioutil.ReadFile(file_names_arr[0])
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading plaintext file.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  key, err := ioutil.ReadFile(file_names_arr[1])
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading key file.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  // Set up variables and remove line feed (ASCII 10) at the end of plaintext and key
  const allowed_char string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
  var cipher string = ""
  plaintext = plaintext[:len(plaintext) - 1]
  key = key[:len(key) - 1]
  var plaintext_len int = len(plaintext)

  // Encrypt plaintext using the key.
  for i := 0; i < plaintext_len; i ++ {
    var encrypt_idx int = 0
    var idx_1 int = 0
    var idx_2 int = 0
    // If the character is a space, then idx_1 is 26 because " " is index 26 of the allowed_char array.
    // If not, % by 65 because "A" in ASCII is 65 and using a % between a character and 65 will result in 
    // the index of the letter in allowed_char.
    if string(plaintext[i]) == " " {
      idx_1 = 26
    } else {
      idx_1 = int(plaintext[i]) % 65
    }
    // Apply the same logic to the key.
    if string(key[i]) == " " {
      idx_2 = 26
    } else {
      idx_2 = int(key[i]) % 65
    }
    // Find the encrypt_idx.
    encrypt_idx = (idx_1 + idx_2) % 27
    // Get the corresponding character from allowed_char and add it to cipher.
    cipher += string(allowed_char[encrypt_idx])
  }
  _, err = conn.Write([]byte(cipher))
  if err != nil {
    fmt.Println("Error sending cipher to the client.")
    fmt.Println(err)
    os.Exit(1)
  }
}


func split_file_names(file_names string) []string {
  file_names_arr := make([]string, 0)
  files := []rune(file_names)
  curr := ""
  for _, c := range files {
    if c != ' ' {
      curr += string(c)
    } else {
      file_names_arr = append(file_names_arr, curr)
      curr = ""
    }
  }
  file_names_arr = append(file_names_arr, curr)
  return file_names_arr
}
