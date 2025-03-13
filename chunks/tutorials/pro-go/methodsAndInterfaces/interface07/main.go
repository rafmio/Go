// Using the empty interface
package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

func main() {
	// var expense Expense = &Product{ "Kayak", "Watersports", 275.00 }
	// The same code below:
	var expense Expense
	var instance *Product = &Product{"Kayak", "Watersports", 275.00}
	expense = instance

	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Product{"Sumbarine", "Watersports", 95000},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:",
				value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)

		}
	}
	fmt.Println()
	fmt.Println(expense)
}
