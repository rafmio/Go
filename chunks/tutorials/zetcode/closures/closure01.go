package main

import "fmt"

type smfnc func(int, int, int) int

func main() {
  sum := func(a, b, c int) int {
    return a + b + c
  } (3, 5, 7)
  fmt.Println("5 + 3 + 7 = ", sum)
  fmt.Printf("Type of sum: %T\n", sum)
  fmt.Println("sum:", sum)

}


// https://zetcode.com/golang/closure/
