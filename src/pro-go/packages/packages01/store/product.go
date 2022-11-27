// Package store provides types and methods
// commonly required for online sales
package store

var standardTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
    Name, Category string // Name and type of the product
    price float64
}

func NewProduct(name, category string, price float64) *Product {
    return &Product{name, category, price}
}

func (p *Product) Price() float64 {
//     return p.price
    return standardTax.calcTax(p)
}

func (p *Product) SetPrice(newPrice float64) {
    p.price = newPrice
}

// The name specified by the package statement should the 
// name of the folder in which the code files are created,
// which is store in this case
