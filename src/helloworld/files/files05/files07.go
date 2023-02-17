package main

import (
  "log"
  "os"
)

func main() {
  file, err := OpenFile("test.txt", os.O_WRONLY, 0666)
  if err != nil {
    if os.IsPermission(err) {
      log.Println("Error: Write primission denied")
    }
  }
  file.Close()

  file, err := os.OpenFile()
}
