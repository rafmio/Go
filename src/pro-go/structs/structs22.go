// Understanding zero value for structs and pointers to structs
package main

import "fmt"

type Product struct {
  name, category string
  price float64
}

func main() {
  var prod Product
  var prodPtr *Product

  fmt.Println("Value:", prod.name, prod.category, prod.price)
  fmt.Println("Pointer:", prodPtr)
  fmt.Println(prod)
}
