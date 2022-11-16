// Checking for items in a map
package main

import (
  "fmt"
  "sort"
)

func main() {
  products := map[string]float64 {
    "Kayak" : 279,
    "Lifejacket" : 48.95,
    "Hat" : 0,
  }
  value, ok := products["Hat"]
  if (ok) {
    fmt.Println("Stored value: ", value)
  } else {
    fmt.Println("No stored value")
  }

  // Using an initialization statement
  if val, okey := products["Nipple"]; okey {
    fmt.Println("Nipple is exists:", val)
  } else {
    fmt.Println("No Nipple in this map")
    fmt.Printf("Type of okey is: %T\n", okey)
  }
  fmt.Println()

  // Deleting items from a map
  delete(products, "Hat")

  if value, ok := products["Hat"]; ok {
    fmt.Println("Stored value:", value)
  } else {
    fmt.Println("No stored value")
    fmt.Println("products:", products)
  }
  fmt.Println()

  for key, value := range products {
    fmt.Println("Key:", key, "Value:", value)
  }
  fmt.Println()

  // Enumerating a map in order
  products2 := map[string]float64 {
    "Kayak" : 279,
    "Lifejacket" : 48.95,
    "Hat" : 0,
    "Anchor" : 12.04,
  }
  keys := make([]string, 0, len(products2))
  for key, _ := range products2 {
    keys = append(keys, key)
  }
  sort.Strings(keys)
  for _, key := range keys {
    fmt.Println("Key:", key, "Value:", products2[key])
  }
}
