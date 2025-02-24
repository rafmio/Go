package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// store a key-value pair
	m.Store("key1", "value1")

	// Load a value by key
	value, ok := m.Load("key1")
	if ok {
		fmt.Println(value)
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v : %v\n", key, value)
		return true
	})
}
