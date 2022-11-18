// Understanding function closure
package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
  fmt.Println("Product:", product, "Price:", calculator(price))
}

func main() {
  watersportsProducts := map[string]float64 {
    "Kayak" : 275,
    "Lifejacket" : 48.95,
  }
  soccerProduct := map[string] float64 {
    "Soccer Ball" : 19.50,
    "Stadium" : 79500,
  }

  calc := func(price float64) float64 {
    if (price > 100) {
      return price + (price * 0.2)
    }
    return price
  }

  for product, price := range watersportsProducts {
    printPrice(product, price, calc)
  }

  calc = func(price float64) float64 {
    if (price > 50) {
      return price + (price * 0.1)
    }
    return price
  }

  for product, price := range soccerProduct {
    printPrice(product, price, calc)
  }
}
