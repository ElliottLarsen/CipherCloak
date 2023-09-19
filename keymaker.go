package main

import (
  "fmt"
  "os"
  "strconv"
  "math/rand"
)

func main() {
  var key_len int
  var err error
  var key string = ""
  // Allowed characters for the key.
  const allowed_char string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
  args := os.Args
  if len(args) == 2 {
    key_len, err = strconv.Atoi(args[1])
    if err != nil {
      panic(err)
    }
  } else {
    fmt.Fprintln(os.Stderr, "Incorrect number of command line arguments.\nExiting...")
    // Program did not crash so exit code is set to 0.
    os.Exit(0)
  }
  // Build the key.
  for i := 0; i < key_len; i++ {
    // Choose a random character from allowed_char and concatenate to return_str.
    key += string(allowed_char[rand.Intn(27)])
  }
  fmt.Fprintln(os.Stdout, key)
}
