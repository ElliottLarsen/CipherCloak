package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  // Check the arguments passed.
  if len(os.Args) != 4 {
    fmt.Fprintln(os.Stderr, "Incorrect command line arguments.  It should be ./client <plaintext> <key> <port>")
    os.Exit(0)
  }

  plaintext := os.Args[1]
  key := os.Args[2]
  port := os.Args[3]
  // Connect to the encrypt server.
  conn, err := net.Dial("tcp", "localhost:" + port)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error connecting to the server on port " + port)
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  defer conn.Close()
  // Send file names to the server.
  file_names := plaintext + " " + key
  _, err = conn.Write([]byte(file_names))
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error sending files to the server.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  // Receive and print the encrypted text from the server.
  buffer := make([]byte, 200000)
  n, err := conn.Read(buffer)
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error receiving encrypted text from the server.")
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  fmt.Println(string(buffer[:n]))

}
