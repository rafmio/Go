// Package store provides types and methods
// commonly required for online salse
package store

var standardTax = newTaxRate(0.25, 20)

//Product describes an item for sale
type Product struct {
    Name, Category string   // Name and type of the product
    price float64           // Price of the product
}

// NewProduct creates new instance of type *Product
func NewProduct(name, category string, price float64) *Product {
    return &Product{name, category, price}
}

// Price() method returns calculated tax using calcTax()
// method - calcTax method for *taxRate instance that
// receives *Product instance and returns floa64
func (p *Product) Price() float64 {
    return standardTax.calcTax(p)
}

// SetPrice() method for *Product instance than receives
// float64 price, and assings price value to *Product's 
// field price
func (p *Product) SetPrice(newPrice float64) {
    p.price = newPrice
}
