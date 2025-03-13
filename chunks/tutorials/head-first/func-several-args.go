package main

import "fmt"

func severalInts(numbers ...int) {
  fmt.Println(numbers)
}

func severalStrings(strings ...string) {
  fmt.Println(strings)
}

func mix(num int, flag bool,  strings ...string) {
  fmt.Println(num, flag, strings)
}

func main() {
  severalInts(1)
  severalInts(1, 2, 3)

  severalStrings("a", "b")
  severalStrings("a", "b", "c", "d", "e")
  severalStrings()

  mix(1, true, "a", "b")
  mix(2, false, "a", "b", "c", "d", "1")
}
