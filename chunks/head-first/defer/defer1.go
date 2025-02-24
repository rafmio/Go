package main

import "fmt"

func Socialize() {
  defer fmt.Println("1: Goodbye!")
  fmt.Println("2: Hello!")
  fmt.Println("3: Nice weather, eh?")
}

func main() {
  defer Socialize()
  fmt.Println("Msg inside main()")
}
