// Understanding pointer and value receivers
package main

import "fmt"

type Product struct {
  name, category string
  price float64
}

func (product Product) printDetails() {
  fmt.Println("Name:", product.name, "Category:", product.category,
  "Price:", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
  if (product.price > threshold) {
    return product.price + (product.price * rate)
  } else {
    return product.price
  }
}

func main() {
  kayak := &Product{ "Kayak", "Watersports", 275.00 }
  ball := Product{ "Soccer Ball", "Soccer", 19.00 }
  kayak.printDetails()
  ball.printDetails()
}


// A method whose receiver is a pointer type can also be invoked through a
// regular value of the underlying type and vise versa
