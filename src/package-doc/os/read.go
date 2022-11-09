package main

import (
  "os"
  "fmt"
  "log"
)

func main() {
  file, err := os.Open("textFileForRead.txt")
  if err != nil {
    log.Fatal(err)
  }

  data := make([]byte, 100)
  count, err := file.Read(data)
  fmt.Printf("Type of count: %T\n", count)
  fmt.Printf("read %d bytes: %q\n", count, data[:count])
  for i, str := range data {
    fmt.Println(i, string(str))
    if i > 62 {
      break
    }
  }
  fmt.Println(data)
  fmt.Println(string(data))
}
