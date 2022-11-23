// Defining an interface
// Using an interface in a function
package main

import "fmt"

type Expense interface {
  getName() string
  getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
  for _, item := range expenses {
    total += item.getCost(true)
  }
  return
}

func main() {
  expenses := []Expense {
    Product { "Kayak", "Watersports", 275.00 },
    Product { "Boat", "Watersports", 195.00 },
    Service { "Boat Cover", 12, 89.50 },
  }

  for _, expense := range expenses {
    fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
  }

  fmt.Println()
  fmt.Println("Total:", calcTotal(expenses))

}


// Variables whose type is an interface have two types:
// - the static type
// - the dynamic type
// The static type is an interface type. The dynamic type is the type of value
// assinged to the variable that implements the interface, such as Product or
// Service in this case
// The static type never changes - the static type of an Expense variable
// is always Expense, but the dynamic type can change by assigning a new value
// of a different type that implements the interface
