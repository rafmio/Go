// Defining variadic parameters
package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
  if (len(suppliers) == 0) {
    fmt.Println("Product:", product, "Supplier: (none)")
  } else {
    for _, supplier := range suppliers {
      fmt.Println("Product:", product, "Supplier:", supplier)
    }
  }
}

func main() {
  printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
  printSuppliers("Lifejacket", "Sail Safe Go")
  printSuppliers("Soccer Ball")
}
