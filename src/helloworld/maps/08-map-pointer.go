package main

import (
	"fmt"
)

func modifyMap(mp map[int]string) {
	// Modify the value of a key
	mp[2] = "two-modified"

	// Add a new key-value pair
	mp[4] = "four"

	// Delete a key-value pair
	delete(mp, 3)
}

func main() {
	mp := make(map[int]string)
	// Initialize the map
	mp[1] = "one"
	mp[2] = "two"
	mp[3] = "three"

	// Iterate over the map and print the key-value pairs
	for k, v := range mp {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}

	// Modify the map using a separate function
	modifyMap(mp)

	// Iterate over the modified map and print the key-value pairs
	fmt.Println("\nModified Map:")
	for k, v := range mp {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}
