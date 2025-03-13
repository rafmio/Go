package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	kayak := Product{
		name:     "Kayak",
		category: "Waterspoerts",
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)
}

// You may see conde that usese the built-in new function to create struct values:
// var lifejacket = new(Product)
// The result is a pointer to a struct value whose fields are initialized with
// their type's zero value. This is equivalent to this statenent:
// var lifejacket = &Product{}
