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
    go decrypt_text(conn)
  }
}

func decrypt_text(conn net.Conn) {
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

  // Open cipher and key files.
  cipher, err := ioutil.ReadFile(file_names_arr[0])
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading cipher file.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  key, err := ioutil.ReadFile(file_names_arr[1])
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading key file.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  // Set up variables and remove line feed (ASCII 10) at the end of plaintext and key.
  const allowed_char string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
  var plaintext string = ""
  cipher = cipher[:len(cipher) - 1]
  key = key[:len(key) - 1]
  var cipher_len int = len(cipher)

  // Decrypt cipher using the key.
  for i := 0; i < cipher_len; i++ {
    var decrypt_idx int = 0
    var idx_1 int = 0
    var idx_2 int = 0
    if string(cipher[i]) == " " {
      idx_1 = 26
    } else {
      idx_1 = int(cipher[i]) % 65
    }

    if string(key[i]) == " " {
      idx_2 = 26
    } else {
      idx_2 = int(key[i]) % 65
    }
    // Find the decrypt_idx.
    decrypt_idx = idx_1 - idx_2
    if decrypt_idx < 0 {
      decrypt_idx += 27
    }
    // Get the corresponding character from allowed_char and add it to plaintext.
    plaintext += string(allowed_char[decrypt_idx])
  }
  _, err = conn.Write([]byte(plaintext))
  if err != nil {
    fmt.Println("Error sending plaintext to the client.")
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
