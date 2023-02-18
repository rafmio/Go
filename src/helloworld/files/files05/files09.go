package main

import (
  "os"
  "log"
  "fmt"
)

func main() {
  err := os.Link("original.txt", "original_also.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("creating sym")
  err = os.Symlink("original.txt", "original_also.txt")
  if err != nil {
    log.Fatal(err)
  }

  fileInfo, err := os.Lstat("original_sym.txt")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Linf info: %+v\n", fileInfo)

  err = os.Lchown("original_sym", os.Getuid(), os.Getgid())
  if err != nil {
    log.Fatal(err)
  }
}
