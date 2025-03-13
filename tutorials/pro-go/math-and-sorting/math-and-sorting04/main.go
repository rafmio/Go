package main

import (
  "math/rand"
  "time"
)

var names = []string {"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {
  rand.Seed(time.Now().UnixNano())

  // Shuffling data
  rand.Shuffle(len(names), func(first, second int) {
    names[first], names[second] = names[second], names[first]
  })

  for i, name := range names {
    Printfln("Index %v: Name: %v", i, name)
  }
  fmt.Println()

  // Sorting data
  
}
