package store

type Product struct {
  Name, Category string
  price float64
}

func NewProduct(name, category string, price float64) *Product {
  return &Product{name, category, price}
}

func(p *Product) Price(taxRate float64) float64 {
  return p.price + (p.price * taxRate)
}


// A common convention is to define a conctructor function whose
// name is New<Type>, such as NewProduct