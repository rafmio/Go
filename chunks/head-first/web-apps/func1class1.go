package main

import "fmt"

func sayHi() {
  fmt.Println("Hi")
}

func sayBye() {
  fmt.Println("Bye")
}

func twice(theFunction func()) {
  theFunction()
  theFunction()
}

func main() {
  var myFunction func()
  myFunction = sayHi
  myFunction()
  sayHi()

  fmt.Printf("Type of myFunction: %T\n", myFunction)
  fmt.Printf("Type of sayHi: %T\n", sayHi)
  fmt.Println()

  twice(sayHi)
  twice(sayBye)
  twice(myFunction)
}
