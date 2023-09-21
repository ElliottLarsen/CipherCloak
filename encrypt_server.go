package main

import (
  "fmt"
  //"io/ioutil"
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
  fmt.Println(file_names_arr)
  // TODO
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
