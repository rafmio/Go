// Performing type assertion
package main

import "fmt"

type Expense interface {
  getName() string
  getCost(annual bool) float64
}

func main() {
  expenses := []Expense {
    Service { "Boat Cover", 12, 89.50, []string{} },
    Service { "Paddle Project", 12, 8, []string{} },
  }

  for _, expense := range expenses {
    s := expense.(Service) // asserts that expense is not nil and value in expense is of type Service
    fmt.Println("Service:", s.description, "Price:",
    s.monthlyFee * float64(s.durationMonths))
  }
}

// go dynamic and static types (methods)
