// Using pointers directly
package main

import "fmt"

type Product struct {
  name, category string
  price float64
}

func calcTax(product *Product) *Product {
  if (product.price > 100) {
    product.price += product.price * 0.2
  }
  return product
}

func main() {
  kayak := calcTax(&Product {
    name: "Kayak",
    category: "Watersports",
    price: 275.00,
  })
  fmt.Println("Name:", kayak.name, "Category:",
  kayak.category, "Price:", kayak.price)

  boat := &Product{
    name: "Boat",
    category: "Watersports",
    price: 314.55,
  }
  fmt.Println("Boat:", boat)
  boat = calcTax(boat)
  fmt.Println("Boat:", boat)
}
