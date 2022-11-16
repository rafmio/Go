// Copying slices with different sizes
package main

import "fmt"

func main() {
  products := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
  replacementProducts := []string { "Canue", "Boots" }
  fmt.Println("len(products):", len(products), "cap(products):", cap(products))
  fmt.Println("len(replacementProducts):", len(replacementProducts),
                "cap(replacementProducts):", cap(replacementProducts))

  fmt.Println()
  copy(products, replacementProducts)
  fmt.Println("products:", products)
  fmt.Println()

  fmt.Println("len(products):", len(products), "cap(products):", cap(products))
  fmt.Println("len(replacementProducts):", len(replacementProducts),
                "cap(replacementProducts):", cap(replacementProducts))
  fmt.Println("--------------------------------------------------------------")
  //
  // products2 := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
  // replacementProducts2 := []string { "Canoe", "Boots" }
  //
  // copy(replacementProducts2, products2)
  //
  // fmt.Println("replacementProducts2:", replacementProducts2)
  //
  // Output is: replacementProducts2: [Kayak Lifejacket]

  products2 := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
  replacementProducts2 := []string { "Canoe", "Boots" }

  copy(products2[0:1], replacementProducts2)

  fmt.Println("products2:", products2)
}
