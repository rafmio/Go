// https://www.geeksforgeeks.org/io-pipe-function-in-golang-with-examples/
package main

import (
  "bytes"
  "fmt"
  "io"
)

func main() {
  pipeReader, pipeWriter := io.Pipe()

  go func() {
    fmt.Fprint(pipeWriter, "Kissy-Missy and Huggy-Wuggy\n")
    pipeWriter.Close()
  }()

  buffer := new(bytes.Buffer)
  buffer.ReadFrom(pipeReader)
  fmt.Println(buffer.String())
}
