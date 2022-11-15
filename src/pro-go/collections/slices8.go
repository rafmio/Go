package main

import "fmt"

func main() {
  products := [4]string { "Kissy", "Missy", "Tosi", "Bosi" }
  allNames := products[1:]
  someNames := []string { "Boots", "Canoe" }
  fmt.Println("someNames:", someNames)
  fmt.Println("allNames:", allNames)
  
  copy(someNames[1:], allNames[2:3])

  fmt.Println("someNames:", someNames)
  fmt.Println("allNames:", allNames)
}
