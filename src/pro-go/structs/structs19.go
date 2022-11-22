// Using pointers types for struct fields
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price - 10, supplier}
}

func newSupplier(name, city string) *Supplier {
	return &Supplier{name, city}
}

func main() {
	acme := &Supplier{"Acme Co", "New York"}
	almazAtn := newSupplier("Almaz Anthey", "Moscow")

	products := [4]*Product{
		newProduct("Kayak", "Watersports", 275.00, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
		newProduct("Boat", "Watersports", 211.00, acme),
		newProduct("Nuclear submarine", "Warship games", 900000, almazAtn),
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}
}
