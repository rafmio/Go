package main

import (
  "fmt"
  "time"
)

type Person struct {
  name string
}

func (p *Person) printName() {
  fmt.Println("Name is: ", p.name)
}

func main() {
  go hello("Martin")
  go hello("Lucia")
  go hello("Mickael")
  go hello("Jozef")
  go hello("Peter")

  var dude1 *Person = &Person {
    name: "Valera",
  }
  
  go dude1.printName()

  time.Sleep(time.Second)


}

func hello(name string) {
  fmt.Printf("Hello %s!\n", name)
}

// https://zetcode.com/golang/goroutine/
