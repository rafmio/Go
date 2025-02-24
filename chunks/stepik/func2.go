package main

import "fmt"

var x = 13

func change() {
  x += 1
  fmt.Println("change():", x)
}

func set(x int) {
  x += 3
  fmt.Println("set():", x)
}

func main() {
  change()
  set(x)
  fmt.Println(x)
}
