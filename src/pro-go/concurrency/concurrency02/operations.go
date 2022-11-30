package main

import (
  "fmt"
  "time"
)

func CalcStoreTotal(data ProductData) {
  var storeTotal float64

  fmt.Println(data)
  fmt.Printf("Type of data is: %T\n", data)
  fmt.Println("data['Chess']:", data["Chess"])
  fmt.Println()

  for category, group := range data {
    // storeTotal += group.TotalPrice(category)
    go group.TotalPrice(category)
  }
  fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, ) (total float64) {
  for _, p := range group {
    fmt.Println(category, "product:", p.Name, "Price:", p.Price)
    total += p.Price
    time.Sleep(time.Millisecond * 100)
  }
  fmt.Println(category, "subtotal:", ToCurrency(total))

  return 
}
