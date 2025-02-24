package main

import (
  "fmt"
  "strconv"
)

func main() {
  val1 := "48.95"
  float1, float1err := strconv.ParseFloat(val1, 64)
  if float1err == nil {
    fmt.Println("Parsed value:", float1)
  } else {
    fmt.Println("Cannot parse", val1, float1err)
  }

  fmt.Println()

  val2 := "4.895e+01"
  float2, float2err := strconv.ParseFloat(val2, 64)
  if float2err == nil {
    fmt.Println("Parsed value:", float2)
  } else {
    fmt.Println("Cannot parse", val2, float2err)
  }
}

// func ParseFloat(s string, bitSize int)(float64, error)
// bitSize:
// 32 - float32
// 64 - float64
// bitSize - size of result
// The result from the functions is always float64, but if 32 is specified,
// then the result can be explicitly converted to a float32
