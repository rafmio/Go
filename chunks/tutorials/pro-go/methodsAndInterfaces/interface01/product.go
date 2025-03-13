package main

type Product struct {
  name, category string
  price float64
}

func (p Product) getName() string {
  return p.name
}

func (p Product) getCost(_ bool) float64 {
  return p.price
}

// To implement an interface, all the methods specified by the interface
// must be defined for a struct type
// Methods must have the same name, paremeter types, and result types
