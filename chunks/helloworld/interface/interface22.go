// Empty interface
// https://www.programiz.com/golang/empty-interface
package main

import "fmt"

func displayValue(i interface{}) {
  fmt.Println(i)
}

func displayAnyNumberValue(i... interface{}) {
  fmt.Println(i)
}

func main() {
  a := "Welcome!"
  b := 20
  c := false

  displayValue(a)
  displayValue(b)
  displayValue(c)

  fmt.Println()

  displayAnyNumberValue(a)
  displayAnyNumberValue(a, b)
  displayAnyNumberValue(a, b, c)
}
