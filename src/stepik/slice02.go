package main

import "fmt"

func main() {
  res := 0
  nums := [3]int{2, 4, 6}

  for i, v := range nums {
    if i % 2 == 0 {
      fmt.Println(i, v, res)
      res += v
    }
  }

  fmt.Println(res)
}
