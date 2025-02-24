package main

import (
	"composition02/store"
	"fmt"
)

func main() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for i, val := range products {
		fmt.Printf("Type of %s is : %T\n", i, val)
	}
	fmt.Println()

	for key, p := range products {
		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	}
	fmt.Println()

	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
	fmt.Println()

	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
			"Price:", item.(store.ItemForSale).Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
	fmt.Println()

	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
			"Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
}
