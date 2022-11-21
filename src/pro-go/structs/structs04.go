// Using field positions to create struct values
package main

import "fmt"

func main() {
  type Product struct {
    name, category string
    price float64
  }

  var kayak = Product{ "Kayak", "Watersports", 275.00 }

  fmt.Println("Name:", kayak.name)
  fmt.Println("Category:", kayak.category)
  fmt.Println("Price:", kayak.price)
}
