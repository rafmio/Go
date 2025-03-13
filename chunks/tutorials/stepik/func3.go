package main

import "fmt"

func double(a int) {
  fmt.Println(a * 2)
}

func main() {
  for i := 4; i > 0; i-- {
    defer double(i)
  }
}
