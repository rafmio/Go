package main

import (
  "fmt"
  "strconv"
)

func main() {
  val1 := "101011"
  int1, int1err := strconv.ParseInt(val1, 2, 8)
  if int1err == nil {
    smallInt := int8(int1)
    fmt.Println("Parsed value:", smallInt)
  } else {
    fmt.Println("Cannot parse", val1, int1err)
  }

  fmt.Println()

  val2 := "0b1100100"
  int2, int2err := strconv.ParseInt(val2, 0, 8)
  if int2err == nil {
    smallInt := int8(int2)
    fmt.Println("Parsed value:", smallInt)
  } else {
    fmt.Println("Cannot parse", val2, int2err)
  }
}
