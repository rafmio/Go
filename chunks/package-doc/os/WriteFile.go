package main

import (
  "log"
  "os"
)

func main() {
  textForWrite := []byte("Hello-Mello, Tosi-Bosi\n")
  err := os.WriteFile("hello.txt", textForWrite, 0666)
  if err != nil {
    log.Fatal(err)
  }
}
