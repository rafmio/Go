// Understanding the effect of pointer method receivers
package main

import "fmt"

type Expense interface {
  getName() string
  getCost(annual bool) float64
}

func main() {
  product := Product { "Kayak", "Watersports", 275.00 }

  var expense Expense = &product

  fmt.Println(product)
  fmt.Println(expense)
  fmt.Println()

  product.price = 100

  fmt.Println(product)
  fmt.Println(expense)
  fmt.Println()

  fmt.Println("Product field value:", product.price)
  fmt.Println("Expense method result", expense.getCost(false))
  fmt.Println("Expense method result", product.getCost(false))
}
