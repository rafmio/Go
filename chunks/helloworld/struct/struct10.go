package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
	Single bool
}

func main() {
	ubay := Person{
		Name:   "John",
		Gender: "Female",
		Age:    17,
		Single: false,
	}

	values := reflect.ValueOf(ubay)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}

	fmt.Println()

	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Name)
	}
}
