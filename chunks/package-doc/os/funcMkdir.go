// https://pkg.go.dev/os
package main

import (
  "log"
  "os"
)

func main() {
  err := os.Mkdir("testDir", 0750)
  if err != nil && !os.IsExist(err) {
    log.Fatal(err)
  }
  //
  // err = os.WriteFile("testDir/testFile.txt", []byte("Hello!"), 0660)
  // if err != nil {
  //   log.Fatal(err)
  // }
}
