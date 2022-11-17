package main

import "fmt"

func printSuppliers(product string, suppliers []string) {
  for _, supplier := range suppliers {
    fmt.Println("Product:", product, "Supplier:", supplier)
  }
}

func main() {
  printSuppliers("Kayak", []string {"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
  printSuppliers("Lifejacket", []string{"Sail Safe Go"})
}
