// Understanding method overloading
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func newSupplier(name, city string) *Supplier {
	return &Supplier{name, city}
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price:", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275.00),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
		p.printDetails()
	}

	suppliers := []*Supplier{
		newSupplier("Acme Co", "New York City"),
		newSupplier("BoatCo", "Chicago"),
	}
	for _, s := range suppliers {
		s.printDetails()
	}
}
