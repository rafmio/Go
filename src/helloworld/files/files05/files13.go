// https://www.devdungeon.com/content/working-files-go
package main

import (
  "os"
  "log"
)

func main() {
  file, err := os.OpenFile(
    "test.txt",
    os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
    0666,
  )
  if err != nil {
    log.Fatal(err, "opening file O_WRONLY | O_TRUNC | O_CREATE")
  }
  defer file.Close()

  byteSlice := []byte("Bytes, bytes and one more byte\n")
  bytesWritten, err := file.Write(byteSlice)
  if err != nil {
    log.Fatal(err, "writing data to file")
  }

  err = file.Sync()
  if err != nil {
    log.Fatal(err, "sync data")
  }

  log.Printf("Wrote %d bytes\n", bytesWritten)
}
