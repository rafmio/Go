// https://golangbyexample.com/interface-in-golang/
// Type switch
package main

import (
	"fmt"
	"os"
)

type animal interface {
	breathe()
	walk()
}

// implementing the lion struct
type lion struct {
	age     int
	hasMane bool
}

func (l *lion) breathe() {
	fmt.Println("Roars")
}

func (l *lion) walk() {
	fmt.Println("Lion walks")
}

// implementing the elephant struct
type elephant struct {
	age      int
	hasTrunk bool
}

func (e *elephant) breathe() {
	fmt.Println("Trumpets")
}

func (e *elephant) walk() {
	fmt.Println("Elephant walks")
}

// implementing the dog struct
type dog struct {
	age     int
	hasTail bool
}

func (d *dog) breathe() {
	fmt.Println("Barks")
}

func (d *dog) walk() {
	fmt.Println("Dog walks")
}

// --- MAIN ---
func main() {
	// reading os.Args for determine which animal to simulate
	var animalName string = "lion" // default to lion if no argument provided in command line arguments
	if len(os.Args) > 1 {
		animalName = os.Args[1]
		fmt.Println("Simulating animal: ", animalName)
	}

	// creating an animal object based on the input name
	var animal animal
	switch animalName {
	case "lion":
		animal = &lion{age: 10, hasMane: true}
	case "elephant":
		animal = &elephant{age: 20, hasTrunk: true}
	case "dog":
		animal = &dog{age: 5, hasTail: true}
	default:
		fmt.Println("Invalid animal name")
		return
	}

	printAnimalInfo(animal)
}

func printAnimalInfo(animal animal) {
	switch v := animal.(type) {
	case *lion:
		fmt.Printf("Lion: Age %d, has mane: %t\n", v.age, v.hasMane)
	case *elephant:
		fmt.Printf("Elephant: Age %d, has trunk: %t\n", v.age, v.hasTrunk)
	case *dog:
		fmt.Printf("Dog: Age %d, has tail: %t\n", v.age, v.hasTail)
	default:
		fmt.Println("Unknown animal type")
	}
	animal.breathe()
	animal.walk()
}
