// Comparing interface values
package main

import "fmt"

type Expense interface {
  getName() string
  getCost(annual bool) float64
}

func main() {
  var e1 Expense = &Product { name: "Kayak" }
  var e2 Expense = &Product { name: "Kayak" }
  var e3 Expense = Service { description: "Boat Cover" }
  var e4 Expense = Service { description: "Boat Cover" }

  fmt.Println("e1 == e2:", e1 == e2)
  fmt.Println("e3 == e4:", e3 == e4)

  var v5 Product = Product { name: "Submarine", category: "Watersports" }
  var e5 Expense = &v5
  fmt.Println(e5)
  fmt.Println(e5.getName())
}

// Product methods receives pointers (product.go)
// Service methods receives values (service.go)

// First two Expense values are not equal. That's because the dynamic type of
// these values is a pointer type, and pointers are equal only if they point
// to the same memeory location

// The second two Expense values are equal because the are simple struct values
// with the same field values
