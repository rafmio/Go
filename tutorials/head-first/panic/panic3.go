package main

import "fmt"

func main() {
  fmt.Println("msg in main()")
  one()
}

func one() {
  defer fmt.Println("deffered in one()")
  two()
}

func two() {
  defer fmt.Println("deffered in two()")
  panic("Let's see what's beed deffered!")
}
