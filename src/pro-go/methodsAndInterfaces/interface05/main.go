// Performing type assertion
package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Project", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275.00},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok { // asserts that expense is not nil and value in expense is of type Service
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths), "---", s)
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:",
				expense.getCost(true))
		}
	}
}

// go dynamic and static types (methods)
// Type assertions can be applied only to interfaces
// Type assertions used to tell the compiler than an interface value has
// a specific dynamic type
