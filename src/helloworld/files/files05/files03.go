package main

import (
  "fmt"
  "log"
  "os"
)

var (
  fileInfo os.FileInfo
  err error
)

func main() {
  fileInfo, err := os.Stat("test.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("fileInfo.Name(): ", fileInfo.Name())
  fmt.Println("fileInfo.Size(): ", fileInfo.Size())
  fmt.Println("fileInfo.Mode(): ", fileInfo.Mode())
  fmt.Println("Last modified: ", fileInfo.ModTime())
  fmt.Println("Is Directory: ", fileInfo.IsDir())
  fmt.Printf("System interface type: %T\n", fileInfo.Sys())
  fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

  fmt.Println()
  sysFileInfo := fileInfo.Sys()
  fmt.Printf("Type of sysFileInfo: %T\n", sysFileInfo)
}
