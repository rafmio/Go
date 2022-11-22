// Understanding struct constructor functions
package main

import "fmt"

type Product struct {
  name, category string
  price float64
}

func newProduct(name, category string, price float64) *Product {
  return &Product{name, category, price}
}

func main() {
  products := [3]*Product {
    newProduct("Kayak", "Watersports", 275.00),
    newProduct("Hat", "Skiing", 42.50),
    newProduct("Submarine", "Warship", 100000),
  }

  for _, p := range products {
    fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
  }
  fmt.Println()

  for i, p := range products {
    fmt.Println(i, p)
  }
  fmt.Println()

  fmt.Println(products[0].name)
  fmt.Println(products[1].name)
  fmt.Println(products[2].name)
}

// Constructor functions are used to create struct values consistently.
// Constructor functions are usually named 'new' or 'New' followed by the
// struct type (e.g. 'Product')
