// https://zetcode.com/golang/builder/
package main

import (
  "fmt"
  "strings"
  )

func main() {
  builder := strings.Builder{}

  animals := "hawk"
  n := 3

  builder.WriteString("There are ")
  builder.WriteString(fmt.Sprintf("%d %s ", n, animals))
  builder.WriteString("in the sky")

  msg := builder.String()

  fmt.Println(msg)
  fmt.Println()
  fmt.Println("builder.Len()", builder.Len())
  fmt.Println("builder.Cap()", builder.Cap())

  fmt.Println(builder.String())
}
