package main

import "fmt"

func getProductName(index int) (name string, err error) {
  if (len(Products) > index) {
    name = fmt.Sprintf("Name of product: %v", Products[index].Name)
  } else {
    err = fmt.Errorf("Error for index %v", index)
  }
  return
}

func main() {
  fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
  fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")

  strg, val2 := fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
  fmt.Println(strg)
  fmt.Println(val2)
  fmt.Println("----------------------------------------")

  name, _ := getProductName(1)
  fmt.Println(name)
  _, err := getProductName(10)
  fmt.Println(err.Error())

}

// Verbs:
// %v - displays the default format for the value
// %#v - displays a value in a format that could be used to re-create the
// value in a Go-code file
// %T - verb displays the Go type of a value
