// assertions
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %s (code: %d)", e.Message, e.Code)
}

func doSomethind() error {
	return &MyError{"Error occurred", 500}
}

func main() {
	person := Person{"Alice", 30}

	var p interface{} = person

	realPerson, ok := p.(Person)
	if !ok {
		fmt.Println("Assertion type error")
		return
	} else {
		fmt.Println("Assertion successful")
	}

	fmt.Printf("Name: %s, Age: %d\n", realPerson.Name, realPerson.Age)
	fmt.Printf("the realPerson is of type %T\n", realPerson)

	// ------------------------
	err := doSomethind()
	if myErr, ok := err.(*MyError); ok {
		fmt.Println(myErr)
	} else {
		fmt.Println("Error occurred but not a MyError")
	}

}
