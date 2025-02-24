// Getting the array underlying a slice
package main

import (
  "fmt"
)

func main() {
  p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
  arrayPtr := (*[3]string)(p1)
  array := *arrayPtr

  fmt.Println(array)
  fmt.Printf("Type of array: %T\n", array)
}
