package main

import "fmt"

type inrootdirStuct struct {
	name  string
	city  string
	price float64
}

func newInrootdirStruct(name string, city string, price float64) inrootdirStuct {
	newInstance := new(inrootdirStuct)
	newInstance.name = name
	newInstance.city = city
	newInstance.price = price

	return *newInstance
}

func sayHelloFromRootDir() {
	fmt.Println("hello from root dir")
}
