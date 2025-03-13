package main

import "fmt"

func main() {
  fmt.Println("Call the welcome()")
  welcome("Hobo-Yobo")
}

func welcome(msg string) {
  fmt.Println(msg)
}
