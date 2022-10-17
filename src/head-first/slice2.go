package main

import "fmt"

func main() {
  numbers := make([]float64, 3)
  numbers[0] = 19.7
  numbers[2] = 25.2
  for i, number := range numbers {
    fmt.Println(i, number)
  }

  var letters = []string{"a", "b", "c"}
  for i, letter := range letters {
    fmt.Println(i, letter)
  }

  fmt.Printf("numbers: %T\n", numbers)
  fmt.Printf("letters: %T\n", letters)

  fmt.Println()

  underlyingArray := [5]string{"a", "b", "c", "d", "e"}
  slice1 := underlyingArray[0:3]
  fmt.Println(slice1)

  i, j := 1, 4
  slice2 := underlyingArray[i:j]
  fmt.Println(slice2)

  slice3 := underlyingArray[2:len(underlyingArray)]
  fmt.Println(slice3)

  slice4 := underlyingArray[:3]
  fmt.Println(slice4)
}
