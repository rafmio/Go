package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create a new person
	p := new(Person)
	p.Name = "John Mohn"
	p.Age = 40
	// p := Person{"John Doe", 30}

	// Marshal the person to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))

	// Unmarshal the JSON data back into a person
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the unmarshalled person
	fmt.Printf("Name: %s, Age: %d\n", newPerson.Name, newPerson.Age)
}
