package main

import "fmt"

func callFunction(passedFunction func()) {
  passedFunction()
}

func callTwice(passedFunction func()) {
  passedFunction()
  passedFunction()
}

func callWithArguments(passedFunction func()) {
  passedFunction("This sentence is", false)
}

func printReturnValue(passedFunction func() string) {
  fmt.Println(???)
}

func functionA() {
  fmt.Println("function called")
}
func functionB() string {
  fmt.Println("function called")
  return "Returning from function"
}

func functionC(a string, b bool) {
  fmt.Println("function called")
  fmt.Println(a, b)
}

func main() {
  callFunction(???)
  callTwice(???)
  callWithArguments(functionC)
  printReturnValue(functionB)
}
