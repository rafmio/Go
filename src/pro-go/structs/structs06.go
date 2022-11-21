// Defining an additional field
package main

import "fmt"

func main() {
  type Product struct {
    name, category string
    price float64
  }

  type StockLevel struct {
    Product
    Alternative Product
    count int
  }

  stockItem := StockLevel {
    Product : Product { "Kayak", "Watersports", 275.00 },
    Alternative: Product { "Lifejacekt", "Watersports", 48.95 },
    count: 100,
  }

  fmt.Println("Name: ", stockItem.Product.name)
  fmt.Println("Category: ", stockItem.Product.category)
  fmt.Println("Price: ", stockItem.Product.price)

  fmt.Println("Alt name: ", stockItem.Alternative.name)
  fmt.Println("Alt category: ", stockItem.Alternative.category)
  fmt.Println("Alt price: ", stockItem.Alternative.price)

  fmt.Println("stockItem count: ", stockItem.count)
}
