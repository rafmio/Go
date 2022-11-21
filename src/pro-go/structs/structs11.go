// Creating arrays, slices, and maps containing struct values
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

  array := [1]StockLevel {
    {
      Product: Product { "Kayak", "Watersports", 275.00 },
      Alternative: Product { "Lifejacket", "Watersports", 48.95 },
      count: 100,
  },
}

  fmt.Println("Array:", array[0].Product.name)

  slice := []StockLevel {
    {
      Product: Product { "Kayak", "Watersports", 275.00 },
      Alternative: Product { "Lifejacket", "Watersports", 48.95 },
      count: 100,
    },
  }
  fmt.Println("Slice, name:", slice[0].Product.name)
  fmt.Println("Slice, category: ", slice[0].Product.category)
  fmt.Println("Slice, price:", slice[0].Product.price)

  kvp := map[string]StockLevel {
    "Kayak" : {
      Product: Product { "Kayak", "Watersports", 275.00 },
      Alternative: Product { "Lifejacket", "Watersports", 48.95 },
      count: 100,
    },
  }
  fmt.Println("Map:", kvp["Kayak"].Product.name)
}
