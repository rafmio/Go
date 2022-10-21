package main

import "fmt"

func main() {
	var testmap map[string]string
	testmap = make(map[string]string)

	testmap["Kissy"] = "Missy"
	value, ok := testmap["Kissy"]
	fmt.Println(value, ok)
}
