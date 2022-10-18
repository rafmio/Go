package main

import "fmt"

func welcome() {
  fmt.Println("Welcome, guys!")
}

func main() {
  defer welcome()
  fmt.Println("Hey!")


}
