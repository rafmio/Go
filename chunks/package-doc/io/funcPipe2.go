// https://golangify.com/io-pipe
package main

import (
  "io"
  "log"
  "os"
  "os/exec"
)

func main() {
  pReader, pWriter := io.Pipe()

  cmd := exec.Command("echo", "Hello Go!\nThis is expample")
  cmd.Stdout = pWriter

  go func() {
    defer pReader.Close()
    if _, err := io.Copy(os.Stdout, pReader); err != nil {
      log.Fatal(err)
    }
  }()

  if err := cmd.Run(); err != nil {
    log.Fatal(err)
  }
}
