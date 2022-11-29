package main

import "fmt"

// Parameter is a slice of *Product
func CalcStoreTotal(data ProductData) {
  var storeTotal float64
  for category, group := range data {
    storeTotal += group.TotalPrice(category)
  }
  fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, ) (total float64) {
  for _, p := range group {
    total += p.Price
  }
  fmt.Println(category, "subtotal:", ToCurrency(total))
  return
}


// Methods can be defined only on types that are created in the same package
