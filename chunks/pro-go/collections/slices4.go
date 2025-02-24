package main

import "fmt"

func main() {
  products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
  someNames := products[1:3]
  allNames := products[:]
  fmt.Println("products:", products)
  fmt.Println("products len:", len(products), "cap:", cap(products))
  fmt.Println("someNames:", someNames)
  fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
  fmt.Println("allNames:", allNames)
  fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
  fmt.Printf("products: %p\n", &products)
  fmt.Printf("someNames: %p\n", &someNames)
  fmt.Printf("allNames: %p\n", &allNames)
  fmt.Println()

  someNames = append(someNames, "Gloves")
  fmt.Println("products:", products)
  fmt.Println("products len:", len(products), "cap:", cap(products))
  fmt.Println("someNames:", someNames)
  fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
  fmt.Println("allNames:", allNames)
  fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
  fmt.Printf("products: %p\n", &products)
  fmt.Printf("someNames: %p\n", &someNames)
  fmt.Printf("allNames: %p\n", &allNames)
  fmt.Println()

  someNames = append(someNames, "Boots")
  fmt.Println("products:", products)
  fmt.Println("products len:", len(products), "cap:", cap(products))
  fmt.Println("someNames:", someNames)
  fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
  fmt.Println("allNames:", allNames)
  fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
  fmt.Printf("products: %p\n", &products)
  fmt.Printf("someNames: %p\n", &someNames)
  fmt.Printf("allNames: %p\n", &allNames)
  fmt.Println()

}
