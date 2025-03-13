package main

import (
  // "fmt"
  "math/rand"
  "time"
)

func IntRange(min, max int) int {
  return rand.Intn(max - min) + min
}

func main() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < 5; i++ {
    Printfln("Value %v : %v", i, IntRange(10, 20))
  }

}
