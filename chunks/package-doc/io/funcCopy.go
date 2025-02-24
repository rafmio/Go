// Copy
// func Copy(dst Writer, src Reader) (written int64, err error)
package main
import (
  "io"
  "log"
  "os"
  "strings"
  "fmt"
)

func main() {
  r := strings.NewReader("some io.Reader stream to be read\n")

  if _, err := io.Copy(os.Stdout, r); err != nil {
    log.Fatal(err)
  }

  io.Copy(os.Stdout, strings.NewReader("Me\n"))
  res, err := io.Copy(os.Stdout, strings.NewReader("Me\n"))
  fmt.Println(res, err)
}

// https://pkg.go.dev/io
// A successful Copy returns err == nil, not err == EOF
