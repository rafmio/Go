package main

import (
	"fmt"
	"math/rand"
)

func returnRandomDataType() interface{} {
	switch rand.Intn(5) {
	case 0:
		return 42
	case 1:
		return 3.14
	case 2:
		return true
	case 3:
		return "Hello, World!"
	case 4:
		return []int{1, 2, 3, 4, 5}
	default:
		return nil
	}
}

func main() {
	value := returnRandomDataType()
	switch v := value.(type) {
	case int:
		println("The value is an integer:", v)
	case float64:
		println("The value is a floating-point number:", v)
	case bool:
		println("The value is a boolean:", v)
	case string:
		println("The value is a string:", v)
	case []int:
		println("The value is an array:", v)
	default:
		println("The value is of unknown type")
	}

	fmt.Println(value)
}
