package main

import "fmt"

type creature struct {
	creatureType string
	amphibian    bool
}

type man struct {
	creature
	gender  string
	age     int
	habbits []string
}

type person struct {
	man
	name string
}

func changeFields(pers *person) {
	pers.name = "Barbarossa"
	pers.gender = "female"
}

func main() {
	pers := new(person)
	pers.name = "Henry"
	pers.gender = "male"
	pers.age = 42
	pers.habbits = make([]string, 3)
	pers.habbits[0] = "reading"
	pers.habbits[1] = "gaming"
	pers.habbits[2] = "programming"

	pers.amphibian = false
	pers.creatureType = "human"

	fmt.Println(pers)

	fmt.Printf("Type of pers: %T\n", pers)

	fmt.Printf("gender variable: %v, address: %p\n", pers.gender, &pers.gender)
	fmt.Printf("age variable: %v, address: %p\n", pers.age, &pers.age)
	fmt.Printf("habbits variable: %v, address: %p\n", pers.habbits, &pers.habbits)
	fmt.Printf("name variable: %v, address: %p\n", pers.name, &pers.name)

	for i := 0; i < len(pers.habbits); i++ {
		fmt.Printf("habbits[%d]: %v, address: %p\n", i, pers.habbits[i], &pers.habbits[i])
	}

	fmt.Printf("creatureType variable: %v, address: %p\n", pers.creatureType, &pers.creatureType)

	changeFields(pers)

	fmt.Println(pers.name, pers.gender)
}
