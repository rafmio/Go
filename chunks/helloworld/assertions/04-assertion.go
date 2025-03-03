package main

import "fmt"

func main() {
	mp := make(map[string]any)
	mp["BYTEA"] = make([]byte, 0)
	mp["INET"] = string("194.58.102.129")
	mp["INTEGER"] = int(300)

	fmt.Println(mp)

	mp["NILVAL"] = string("1")

	fmt.Println(mp)

	for key, val := range mp {
		fmt.Printf("key: %s: val: %v, value's type: %T\n", key, val, val)
	}

	fmt.Println()

	for key, val := range mp {
		switch v := val.(type) {
		case string:
			fmt.Printf("key: %s: val: %s, value's type: %T\n", key, v, v)
		case []byte:
			fmt.Printf("key: %s: val: %x, value's type: %T\n", key, v, v)
		case int:
			fmt.Printf("key: %s: val: %d, value's type: %T\n", key, v, v)
		default:
			fmt.Printf("key: %s: val: %v, value's type: %T\n", key, v, v)
		}
	}
	fmt.Println()

	intVal := mp["INTEGER"]
	fmt.Printf("intVal: %d, type: %T\n", intVal, intVal)
	byteArr := mp["BYTEA"]
	fmt.Printf("byteArr: %x, type: %T\n", byteArr, byteArr)
	inetVal := mp["INET"]
	fmt.Printf("inetVal: %s, type: %T\n", inetVal, inetVal)
	nilVal := mp["NILVAL"]
	fmt.Printf("nilVal: %v, type: %T\n", nilVal, nilVal)
}
