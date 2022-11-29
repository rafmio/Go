package store

type SpecialDeal struct {
  Name string
  *Product
  price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal {
  return &SpecialDeal{name, p, p.price - discount}
}

// Add own Price() method for SpecialDeal
func(deal *SpecialDeal) Price(taxRate float64) float64 {
  return deal.price
}
