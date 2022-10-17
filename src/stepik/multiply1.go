package  main

import "fmt"

func main() {
  var number int
  var mult int = 1
  fmt.Scanln(&number)


  for i := 1; i <= number; i++ {
    mult = mult * i
    fmt.Println("i = ", i, "mult = ", mult)
  }
  fmt.Println("mult = ", mult)
}
