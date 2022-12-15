// Empty interface
// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type human struct {
  age int
  name string
}

func main() {
  var man human = human{age: 22, name: "Yorkey"}
  test(man)
  test("this is string")
  test("10")
  test(10)
  test(true)
}

func test(a interface{}) {
  fmt.Printf("%v, %T\n", a, a)
}
