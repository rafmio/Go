package main

import "fmt"

func main() {
  var names [3]string

  names[0] = "Kayak"
  names[1] = "Lifejacket"
  names[2] = "Paddle"
  fmt.Println(names)

  names2 := [...]string{"Kayak", "Lifejacket", "Pallde"}
  fmt.Println(names2)

  otherArray := names
  names[0] = "Canoe"
  fmt.Println("names:", names)
  fmt.Println("otherArray:", otherArray)

  otherArray2 := &names
  names[1] = "Boat"
  fmt.Println("names:      ", names)
  fmt.Println("otherArray2:", otherArray2)
  fmt.Println("*otherArray2:", *otherArray2)
  fmt.Println()

  names3 := [3]string{ "Kayak", "Lifejacket", "Paddle" }
  moreNames3 := [3]string {"Kayak", "Lifejacket", "Paddle"}
  same3 := names3 == moreNames3
  fmt.Println("comparison:", same3)

  for index, value := range names3 {
    fmt.Println("Index:", index, "Value:", value)
  }
  for _, value := range names3 {
    fmt.Println("Value:", value)
  }
}
