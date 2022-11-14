package main

import (
  "fmt"
  "strconv"
)

func main() {
  val1 := true
  val2 := false
  str1 := strconv.FormatBool(val1)
  str2 := strconv.FormatBool(val2)
  fmt.Println("Formatted value 1:" + str1)
  fmt.Println("Formatted value 2:" + str2)
  fmt.Println()

  val3 := 275
  base10String := strconv.FormatInt(int64(val3), 10)
  base2String := strconv.FormatInt(int64(val3), 2)
  fmt.Println("Base 10: " + base10String)
  fmt.Println("Base  2: " + base2String)
  fmt.Println()

  val4 := 336
  base10String4 := strconv.Itoa(val4)
  base2String4 := strconv.FormatInt(int64(val4), 2)
  fmt.Println("Base 10: " + base10String4)
  fmt.Println("Base  2: " + base2String4)
  fmt.Println()

  val5 := 49.95
  Fstring := strconv.FormatFloat(val5, 'f', 2, 64)
  Estring := strconv.FormatFloat(val5, 'e', -1, 64)
  fmt.Println("Format F: " + Fstring)
  fmt.Println("Format E: " + Estring)
  fmt.Println()

}

// func FormatBool(b bool) string
// func FormatInt(i int64, base int) string
// FormatInt accepts only int64 values
//
// func Itoa(i int) string
// Itoa is equivalent to FormatInt(int64(i), 10)
//
// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
// prec (precision) - specifies the number of digits that will
// follow the decimal point
