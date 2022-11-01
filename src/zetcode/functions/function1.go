package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func sub(x, y int) int {
  return x - y
}

func main() {
  fmt.Println(add(10, 8))
  fmt.Println(sub(10, 8))
}

// https://zetcode.com/golang/function/
