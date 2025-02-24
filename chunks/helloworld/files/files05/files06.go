package main

import (
  "log"
  "os"
  "fmt"
)

var (
  fileInfo *os.FileInfo
  err error
  fileName string
)

func main() {
  fmt.Printf("Enter file name: ")
  fmt.Scanf("%s\n", &fileName)

  fileInfo, err := os.Stat(fileName)
  if err != nil {
    if os.IsNotExist(err) {
      log.Fatal("File does not exist")
    }
  }

  log.Println("File does exist. File info:")
  log.Println(fileInfo)
}
