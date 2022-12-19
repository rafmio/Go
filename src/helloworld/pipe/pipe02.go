// https://www.geeksforgeeks.org/io-pipe-function-in-golang-with-examples/
package main

import (
  "bytes"
  "fmt"
  "io"
)

func main() {
  pipeReader, pipeWriter := io.Pipe()
  fmt.Printf("Type of pipeReader: %T, type of pipeWriter: %T\n", pipeReader, pipeWriter)

  go func() {
    fmt.Fprint(pipeWriter, "Geeks\n")
    pipeWriter.Close()
  }()

  buffer := new(bytes.Buffer)
  fmt.Printf("Type of buffer: %T\n", buffer)
  fmt.Println()

  buffer.ReadFrom(pipeReader)

  fmt.Print(buffer.String())

  fmt.Println()
}
