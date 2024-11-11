// JSON tags
// https://golang.cafe/blog/golang-json-marshal-example.html
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	lastLoginAt string
}

func main() {
	u, err := json.Marshal(
		User{
			Name:        "John Doe",
			Age:         30,
			Active:      true,
			lastLoginAt: "2022-01-01T12:00:00Z",
		},
	)
	if err != nil {
		log.Println("Marshalling error:", err)
		return
	}

	fmt.Println(string(u))

	// write to a file
	err = os.WriteFile("03-result.json", u, 0644)
	if err != nil {
		log.Println("Writing file error:", err)
		return
	}
}
