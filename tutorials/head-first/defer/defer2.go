package main

import (
  "fmt"
  "log"
)

func Socialize() error {
  defer fmt.Println("1: Goodbye")
  fmt.Println("2: Hello")
  return fmt.Errorf("3: I don't want to talk")
  fmt.Println("4: Nice weather, eh?")
  return nil
}

func main() {
  err := Socialize()
  if err != nil {
    log.Fatal(err)
  }
    
}
