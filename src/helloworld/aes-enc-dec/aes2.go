package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "fmt"
  "io"
)

func main() {
  text := []byte("Mu super secret code stuff")
  key := []byte("passphrasewhichneedstobe32bytes")

  c, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println(err)
  }


}
