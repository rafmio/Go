// Understanding structs and pointers
package main

import "fmt"

type Product struct {
  name, category string
  price float64
}

func main() {
  p1 := Product {
    name: "Kayak",
    category: "Watersports",
    price: 275.00,
  }

  p2 := &p1

  p1.name = "Original Kayak"

  fmt.Println("P1:", p1.name)
  fmt.Println("P2:", (*p2).name)

  (*p2).category = "New Watersports"
  fmt.Println()
  fmt.Println(p1)
  fmt.Println(*p2)

}
