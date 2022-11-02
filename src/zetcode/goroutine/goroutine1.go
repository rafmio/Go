package main

import (
  "fmt"
  "time"
)

func main() {
  go hello("Martin")
  go hello("Lucia")
  go hello("Mickael")
  go hello("Jozef")
  go hello("Peter")

  time.Sleep(time.Second)
}

func hello(name string) {
  fmt.Printf("Hello %s!\n", name)
}

// https://zetcode.com/golang/goroutine/
