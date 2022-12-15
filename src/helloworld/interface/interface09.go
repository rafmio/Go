// https://golangify.com/parsing-csv
package main

import (
  "fmt"
  "strings"
)

type martian struct{}

func(m martian) talk() string {
  return "nack nack"
}

type laser int

func(l laser) talk() string {
  return strings.Repeat("pew", int(l))
}

var talker interface {
  talk() string
}



func main() {
  talker = martian{}
  fmt.Println(talker.talk())

  talker = laser(3)
  fmt.Println(talker.talk())

  fmt.Printf("Type of talker: %T\n", talker)
  fmt.Println()

}
