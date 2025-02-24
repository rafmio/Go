package main

import "fmt"

func main() {
  a := make([]int, 5)
  a = append(a, 8)
  fmt.Println(a)
  a = append(a, 7)
  fmt.Println(a)

  b := make([]int, 5)
  a[1] = 2
  a[2] = 3

  for i, v := range b {
    fmt.Println(i, v)
  }

  fmt.Println()

  x := "hello"
  for _, c := range x {
    fmt.Println(c)
  }

  fmt.Println()

  y := "hello"
  for _, c := range y {
    fmt.Println(string(c))
  }
}
