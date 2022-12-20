package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, _ := json.Marshal(28192)
	fmt.Println(string(a))
	fmt.Printf("Type of a: %T\n", a)

	b, _ := json.Marshal(true)
	fmt.Println(string(b))
	fmt.Printf("Type of a: %T\n", a)

	sls, _ := json.Marshal([]string{"foo", "bar", "baz"})
	fmt.Println(string(sls))
	fmt.Printf("Type of a: %T\n", a)

	mp, _ := json.Marshal(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	fmt.Println(string(mp))
	fmt.Printf("Type of a: %T\n", a)

}

// func Marshal(v any) ([]byte, error)
