package main

import "fmt"

// Parameter is a map of slices of *Product type
func CalcStoreTotal(data ProductData) {
  var storeTotal float64
  
  fmt.Println(data)
  fmt.Printf("Type of data is: %T\n", data)
  fmt.Println( data["Chess"])
  fmt.Println()

  for category, group := range data {
    // storeTotal += group.TotalPrice(category)
    go group.TotalPrice(category) // Total: $0.00
  }
  fmt.Println("Total:", ToCurrency(storeTotal))
}

// Parameter is a slice of *Product
func (group ProductGroup) TotalPrice(category string, ) (total float64) {
  for _, p := range group {
    fmt.Println(category, "product:", p.Name, "Price:", p.Price)
    total += p.Price
  }
  fmt.Println(category, "subtotal:", ToCurrency(total))
  return
}


// Methods can be defined only on types that are created in the same package
