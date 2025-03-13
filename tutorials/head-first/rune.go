package main

import "fmt"

func main() {
  runestr := []rune("Hello, I'm rune!")
  for i, onerune := range runestr {
    fmt.Printf("%2d - %3d - %s\n", i, onerune, string(onerune))
  }
  fmt.Println("len(runestr): ", len(runestr))
}
