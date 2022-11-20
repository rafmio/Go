// Closures
package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
    fmt.Println("Product:", product, "Price:", calculator(price)
}

var prizeGiveaway = false

func priceCalcFactory(threshold, rate float64, zeroPrices bool) calcFunc {
    return func(price float64) float64 {
        if(zeroPrices) {
            return 0
        }
    }
}
