package main

import (
  "fmt"
  "strconv"
)

func main() {
  val1 := "100"
  int1, int1err := strconv.ParseInt(val1, 10, 0)
  if int1err == nil {
    var intResult int = int(int1)
    fmt.Println("Parsed value:", intResult)
  } else {
    fmt.Println("Cannot parse", val1, int1err)
  }

  fmt.Println()

  val2 := "100"
  int2, int2err := strconv.Atoi(val2)
  if int2err == nil {
    var intResult int = int2
    fmt.Println("Parsed value:", intResult)
  } else {
    fmt.Println("Cannot parse", val2, int2err)
  }
}

// func Atoi(s string)(int, error)
// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int
