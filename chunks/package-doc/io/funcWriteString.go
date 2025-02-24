// https://pkg.go.dev/io
// func WriteString(w Writer, s string) (n int, err error)
package main

import (
  "io"
  "log"
  "os"
  "strings"
)

func main() {
  if _, err := io.WriteString(os.Stdout, "Hello-mello Kissy-Missy\n"); err != nil {
    log.Fatal(err)
  }

  txt := strings.NewReader("Hello-mello Kissy-Missy\n")
  if _, err := io.Copy(os.Stdout, txt); err != nil {
    log.Fatal(err)
  }
}
