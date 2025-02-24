package main

import (
  "os"
  "fmt"
  "log"
)

func main() {
  file, _ := os.Open("test.txt")
  defer file.Close()

  var offset int64 = 5
  var whence int = 0

  newPosition, err := file.Seek(offset, whence)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Just moved to 5: ", newPosition)

  newPosition, err = file.Seek(-2, 1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Just moved back two: ", newPosition)

  currntPosition, err := file.Seek(0, 1)
  fmt.Println("Current position: ", currntPosition)

  newPosition, err = file.Seek(0, 0)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Position after seeking 0, 0: ", newPosition)
}
