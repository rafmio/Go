// Using the empty interface for function parameters
package main

import "fmt"

type Expense interface {
  getName() string
  getCost(annual bool) float64
  getCostAndPrice(annual bool)
  getPerson()
}


func processItem(item interface{}) {
  switch value := item.(type) {
  case Product:
    value.getCostAndPrice(true)
  case *Product:
    value.getCostAndPrice(true)
  case Service:
    value.getCostAndPrice(true)
  case Person:
    value.getPerson()
  case *Person:
    value.getPerson()
  case string, bool, int:
    fmt.Println("Built-in type:", value)
  default:
    fmt.Println("Default:", value)

  }
}

func main() {
  var expense Expense = &Product {"Kayak", "Watersports", 275.00}
  var person1 Expense = Person {"Alice", "London"}
  var person2 Expense = &Person {"Bob", "New York"}

  data := []interface {} {
    expense,
    Product {"Lifejacket", "Watersports", 48.95},
    Service {"Boat Cover", 12, 89.50, []string{}},
    person1,
    person2,
    "This is string",
    100,
    true,
  }

  for _, item := range data {
    processItem(item)
  }
}
