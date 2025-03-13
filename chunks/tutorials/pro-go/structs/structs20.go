// Understanding pointer field copying
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

func main() {
	acme := &Supplier{"Acme Co", "New York"}

	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := *p1		// Shadow copy (неглубокая копия), where pointers are copied,
							// but not the values to which they point

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

	fmt.Println()
	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"

	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}
	fmt.Println()

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}
