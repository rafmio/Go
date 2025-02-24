package main

import (
  "os"
  "log"
  "io"
)

func main() {
  orifinalFile, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err, "opening file")
  }
  defer orifinalFile.Close()

  newFile, err := os.Create("test_copy.txt")
  if err != nil {
    log.Fatal(err, "creating file")
  }
  defer newFile.Close()

  byteWritten, err := io.Copy(newFile, orifinalFile)
  if err != nil {
    log.Fatal(err, "copying file")
  }

  log.Printf("Copied %d bytes", byteWritten)

  err = newFile.Sync()
  if err != nil {
    log.Fatal(err, "sync files")
  }
}
