package main

import "fmt"

func main() {
  fmt.Println("Beginning")

  for i := 0; i < 5; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("End")
}
