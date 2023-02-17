package main

import (
  "log"
  "os"
  "fmt"
)

func main() {
  var file *os.File
  file, err := os.Open("test.txt")
  if err != nil {
    fmt.Println("opening the file: ", err.Error())
    fmt.Println("Creating the file...")
    file, err = os.Create("test.txt")
    if err != nil {
      log.Fatal(err)
    }
  }

  file.Close()

  file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("File was opened")
  }

  file.Close()
}
