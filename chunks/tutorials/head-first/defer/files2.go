package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "path/filepath"
  "os"              
)

func scanDirectory(path string) error {
  fmt.Println(path)
  files, err := ioutil.ReadDir(path)
  if err != nil {
    return err
  }

  for _, file := range files {
    filePath := filepath.Join(path, file.Name())
    if file.IsDir() {
      err := scanDirectory(filePath)
      if err != nil {
        return err
      }
    } else {
      fmt.Println(filePath)
    }
  }
  return nil
}

func main() {
  userPath := os.Args[1]
  err := scanDirectory(userPath)
  if err != nil {
    log.Fatal(err)
  }
}
