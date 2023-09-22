package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  // Check the arguments passed.
  if len(os.Args) != 4 {
    fmt.Fprintln(os.Stderr, "Incorrect command line arguments.  It should be ./client <cipher> <key> <port>")
    os.Exit(0)
  }

  cipher := os.Args[1]
  key := os.Args[2]
  port := os.Args[3]
  // Connect to the decrypt server.
  conn, err := net.Dial("tcp", "localhost:" + port)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error connecting to the decrypt server on port " + port)
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  defer conn.Close()
  // Send file names to the server.
  file_names := cipher + " " + key
  _, err = conn.Write([]byte(file_names))
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error sending files to the decrypt server.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  // Receive and print the decrypted message from the server.
  buffer := make([]byte, 200000)
  n, err := conn.Read(buffer)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error receiving decrypted message from the decrypt server.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  fmt.Println(string(buffer[:n]))
}
