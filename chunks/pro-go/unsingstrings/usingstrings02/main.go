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
  fmt.Println()

  number := 250

  Printfln("Binary: %b", number)
  Printfln("Decimal: %d", number)
  Printfln("Octal: %o, %O", number, number)
  Printfln("Hexadecimal: %x, %X", number, number)
  fmt.Println()

  number2 := 279.00
  Printfln("Decimalles with exponent: %b", number2)
  Printfln("Decimal with exponent: %e", number2)
  Printfln("Decimal without exponent: %f", number2)
  Printfln("Hexadecimal: %x, %X", number2, number2)
  fmt.Println()

  Printfln("Decimal without exponent: >>%8.2f<<", number2)
  Printfln("Decimal without exponent: >>%.2f<<", number2)
  fmt.Println()

  Printfln("Sign: >>%+.2f<<", number2)
  Printfln("Zeros for Padding: >>%010.2f<<", number2)
  Printfln("Right Padding: >>%-8.2f<<", number2)
  fmt.Println()

  name := "Kayak"
  Printfln("String: %s", name)
  Printfln("Character: %c", []rune(name)[0])
  Printfln("Unicode: %U", []rune(name)[0])
  fmt.Println()

  // Boolean
  Printfln("Bool: %t", len(name) > 1)
  Printfln("Bool: %t", len(name) > 100)
  fmt.Println()

  // Pointer formatting verb (%p - hexadecimal representation of pointers)
  Printfln("Pointer: %p", &name)
  Printfln("Pointer: %x", &name)
  fmt.Println()

}
