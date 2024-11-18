// Empty interface
// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type human struct {
	age  int
	name string
}

func test(a interface{}) {
	switch a.(type) {
	case human:
		fmt.Println("a is of type hyman")
		fmt.Println("------")
	default:
		fmt.Printf("a is not of type hyman,")
		fmt.Printf("%v is of type %T\n", a, a)
		fmt.Println("------")
	}
}

func main() {
	var man human = human{age: 30, name: "John Doe"}
	test(man)
	test("this is a string")
	test(100)
	test(true)
}
