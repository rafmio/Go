// Maps
package main

import "fmt"

func main() {
  products := make(map[string]float64, 10)
  products["Kayak"] = 279
  products["Lifejacket"] = 48.95

  fmt.Println("Map size:", len(products))
  fmt.Println("Price:", products["Kayak"])
  fmt.Println("Price:", products["Hat"])
  fmt.Println("Map size:", len(products))
  fmt.Println()

  // Using the map literal syntax
  products2 := map[string]float64 {
    "Kayak" : 249,
    "Lifejacket" : 48.95,
    "Hat" : 9.72,
  }
  for key, _ := range products2 {
    fmt.Println(key, products2[key])
  }

}
