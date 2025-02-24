package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Human struct {
	name   string
	age    int
	gender string
}

func (h *Human) generateAge() {
	h.age = rand.Intn(80)
}

func (h *Human) generateGender() {
	randGenNum := rand.Intn(2)
	if randGenNum == 0 {
		h.gender = "male"
	} else {
		h.gender = "female"
	}
}

func generateNames() []string {
	names := []string{
		"John",
		"Jerry",
		"Sarah",
		"Bill",
		"Abraham",
		"Linda",
		"Katy",
		"Bill",
		"Jassy",
	}

	return names
}

func main() {
	names := generateNames()
	humanList := make([]*Human, 0)

	humanChan := make(chan Human, 2)

	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			person := new(Human)
			person.name = name
			person.generateAge()
			person.generateGender()
			humanChan <- *person
		}(name)
	}

	go func() {
		wg.Wait()
		close(humanChan)
	}()

	go func() {
		person := <-humanChan
		humanList = append(humanList, &person)
	}()

	for i, val := range humanList {
		fmt.Println(i, ";", val.name, ";", val.age, ";", val.gender)
	}
}
