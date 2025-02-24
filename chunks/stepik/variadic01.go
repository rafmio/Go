package main

import "fmt"

func sum(nums ...int) {
  total := 0
  for _, v := range nums {
    total += v
  }
  fmt.Println(total)
}

func main() {
  sum(2, 4, 6)
  sum(10, 20, 30, 40)
  sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

  s := []int{42, 44, 32, 22, 11}
  sum(s...)
}
