package main

import "fmt"

func main() {
  var num int
  var str string

  fmt.Printf("Enter num: ")
  fmt.Scanf("%d", &num)

  fmt.Printf("Enter string: ")
  fmt.Scanln(&str)

  fmt.Println("You entered:")
  fmt.Println(num, str)

}
