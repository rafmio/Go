// Pipe
// func Pipe() (*PipeReader, *PipeWriter)

package main

import (
  "fmt"
  "io"
  "log"
  "os"
)

func main() {
  r, w := io.Pipe()

  go func() {
    fmt.Fprint(w, "Hello-mello there\n", 1000, "\n")
    w.Close()
  }()

  _, err := io.Copy(os.Stdout, r)

  if err != nil {
    log.Fatal(err)
  }

}

// Pipe is a form of redirection from one process to another process
// It is unidirectional (однонаправленный) data channel that can be
// used for interprocess communication (IPC)
