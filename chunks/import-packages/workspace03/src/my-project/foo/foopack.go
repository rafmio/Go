package foo

import "fmt"

type TstStruct struct {
	name string
	age  int
}

var TstString string = "string from foopack"

func TryToImport() {
	tst := TstStruct{
		name: "foo",
		age:  10,
	}

	fmt.Println("from foopack:", tst.name)
	fmt.Println("from foopack:", tst.age)
}
