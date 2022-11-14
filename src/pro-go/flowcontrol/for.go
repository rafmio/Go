package main

import "fmt"

func main() {
  counter := 0
  for {
    fmt.Println("Counter:", counter)
    counter++
    if (counter > 3) {
      break
    }
  }
  fmt.Println("Current counter's value: ", counter)
  fmt.Println()

  for counter2 := 0; true; counter2++ {
    fmt.Println("Counter2:", counter2)
    if (counter2 > 3) {
      break
    }
  }
}
