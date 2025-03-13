package main

import "fmt"

func welcome(name string) {
  fmt.Println("Hello" + name)
}

func main() {
  welcome("Walter")
  welcome("Hank")
}
