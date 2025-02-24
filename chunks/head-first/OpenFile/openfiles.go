package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
  check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text()) // выводит все строки из файла
  }
  check(scanner.Err())
}
