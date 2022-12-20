// JSON tags
// https://golang.cafe/blog/golang-json-marshal-example.html
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	lastLoginAt string
}

func main() {
	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(u))
}
