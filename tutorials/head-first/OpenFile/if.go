package main

import (
  "os"
  "log"
  // "io/ioutil"
)

func saveString(fileName string, str string) error {
  err := os.WriteFile(fileName, []byte(str), 0600)
  return err
}

func main() {
  if err := saveString("english.txt", "Hello\n"); err != nil {
    log.Fatal(err)
  }
  if err := saveString("hindi.txt", "Namaste\n"); err != nil {
    log.Fatal(err)
  }
}

// В учебнике для WriteFile вместо os применялась io/ioutil
// Но в документации сказано:
// Deprecated: As of Go 1.16, this function simply calls os.WriteFile
// т.е. конструкция ioutil.WriteFile устарела, нужно применять os.WriteFile
