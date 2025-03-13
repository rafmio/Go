package main

import (
  "fmt"
  "strconv"
)

func main() {
  val1 := "100"
  int1, int1err := strconv.ParseInt(val1, 0, 8)
  if int1err == nil {
    fmt.Println("Parsed value:", int1)
  } else {
    fmt.Println("Cannot parse", val1)
  }

  val2 := "500"
  int2, int2err := strconv.ParseInt(val2, 0, 8)
  if int2err == nil {
    fmt.Println("Parsed value:", int2)
  } else {
    fmt.Println("Cannot parse", val2, int2err)
  }

  val3 := "100"
  int3, int3err := strconv.ParseInt(val3, 0, 8)
  if int3err == nil {
    smallInt := int8(int3)
    fmt.Println("Parsed value:", smallInt)
  } else {
    fmt.Println("Connot parse", val3, int3err)
  }
}

// strconv.ParseInt(s string, base int, bitSize int)(i int64, err error)
// s - string to parse
// base - base for the number, or zero to let the function detect the base from
// the string's prefix
// base:
// 2 for "0b"
// 8 for "0" or "0o"
// 16 for "0x"
//
// bitSize: 0, 8, 16, 32, 64 - int, int8, int16, int32, int64
// BUT!!! The function always returns an int64
