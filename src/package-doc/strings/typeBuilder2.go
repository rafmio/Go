// https://zetcode.com/golang/builder/
package main

import (
  "fmt"
  "strings"
)

func main() {
  // builder := strings.Builder{}
  var builder strings.Builder

  builder.WriteString("There")
  builder.WriteString(" are")
  builder.WriteString(" three")
  builder.WriteString(" hawks")
  builder.WriteString(" in the sky")

  fmt.Println(builder.String())

  builder.WriteString(". ")
  builder.WriteString("One more text")

  fmt.Println(builder.String())

  builder2 := strings.Builder{}

  data1 := []byte{72, 101, 108, 108, 111}
  data2 := []byte{32}
  data3 := []byte{116, 104, 101, 114, 101, 33}

  builder2.Write(data1)
  builder2.Write(data2)
  builder2.Write(data3)

  fmt.Println(builder2.String())
}
