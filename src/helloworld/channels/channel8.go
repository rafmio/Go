// https://habr.com/ru/post/490336/
package main

import "fmt"

func main() {
  c := make(chan int)

  fmt.Printf("Type of `c` is : %T\n", c)
  fmt.Printf("Value of `c` is : %v\n", c)

  var invalue int = 100
  go printIntvalue(c)
  c <- invalue

}

func printIntvalue(c chan int) {
  var outvalue int
  outvalue = <- c
  fmt.Println("outvalue: ", outvalue)

}
