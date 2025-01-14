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
	// reading JSON data from 03-result.json
	file, err := os.ReadFile("03-result.json")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(file, &user)
	if err != nil {
		log.Fatal(err)
	}

	printStruct(&user)

	user.lastLoginAt = "2023-08-15T12:34:56Z"

	printStruct(&user)
}

// print struct fields
func printStruct(u *User) {
	fmt.Printf("%s\n%d\n%t\n%s\n", u.Name, u.Age, u.Active, u.lastLoginAt)
	fmt.Println("-----------------------------")
}
