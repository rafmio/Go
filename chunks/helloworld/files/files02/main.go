// Writing bytes to a file
package main

import (
  "fmt"
  "os"
)

func main() {
  f, err := os.Create("/home/raf-pc/Go/src/helloworld/files/files02/bytes")
  if (err != nil) {
    fmt.Println(err.Error())
  }

  d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
  n2, err := f.Write(d2)
  if (err != nil) {
    fmt.Println(err.Error())
    f.Close()
  }

  fmt.Println(n2, "bytes written successfully")
  err = f.Close()
  if (err != nil){
    fmt.Println(err.Error())
  }
}
