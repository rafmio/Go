package main

import (
  "fmt"
  "io/ioutil"
  "log"
)

func main() {
  files, err := ioutil.ReadDir("/home/raf/go/Go/src/head-first")
  if err != nil {
    log.Fatal(err)
  }
  for _, file := range files {
    if file.IsDir() {
      fmt.Println("Directory:", file.Name())
    } else {
      fmt.Println("File:", file.Name())
    }
  }
  fmt.Printf("Type of files: %T\n", files)
}

// read list of dir
