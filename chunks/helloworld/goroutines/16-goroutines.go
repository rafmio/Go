package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Human struct {
	gender string
	age    int
}

func main() {
	humanSls := make([]Human, 10)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			person := new(Human)
			person.gender = generateGender()
			person.age = generateAge()

			humanSls[i] = *person
		}()
	}
	wg.Wait()

	for k, v := range humanSls {
		fmt.Println(k, "gender:", v.gender, "age:", v.age)
	}
}

func generateGender() string {
	var gender string

	num := rand.Intn(2)

	if num == 0 {
		gender = "male"
	} else {
		gender = "female"
	}

	return gender
}

func generateAge() int {
	return rand.Intn(95)
}
