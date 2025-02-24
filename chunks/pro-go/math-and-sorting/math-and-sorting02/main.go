package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  for i := 0; i < 5; i++ {
    Printfln("Value %v : %v", i, rand.Int())
  }
  fmt.Println()

  rand.Seed(time.Now().UnixNano())

  for i := 0; i < 5; i++ {
    Printfln("Value %v : %v", i, rand.Int())
  }
  fmt.Println()

  for i := 0; i < 5; i++ {
    Printfln("Value %v : %v", i, rand.Intn(40))
  }

}
