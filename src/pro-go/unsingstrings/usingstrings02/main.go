package main

import "fmt"

func Printfln(template string, values ...interface{}) {
  fmt.Printf(template + "\n", values...)
}

func main() {
  Printfln("Value: %v", Kayak)
  Printfln("Go syntax: %#v", Kayak)
  Printfln("Type: %T", Kayak)
  fmt.Println()

  Printfln("Value: %v", Kayak)
  Printfln("Value with fields: %+v", Kayak)
  fmt.Println()

  fmt.Println(Kayak.String())
}
