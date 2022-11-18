package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
  fmt.Println("Product:", product, "Price:", calculator(price))
}

func main() {
  products := map[string]float64 {
    "Kayak" : 275,
    "Boat" : 199,
    "Lifejacket" : 48.95,
  }
  for product, price := range products {
    printPrice(product, price, func (price float64) float64 {
      return price + (price * 0.2)
    })
  }
}
