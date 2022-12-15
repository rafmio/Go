package main

import "fmt"

type employee interface {
	typing()
	walking()
}

type finansist interface {
	employee
	accrue()
	calculate()
}

type accountant struct {
	experience int
	age        int
}

func (a accountant) typing() {
	fmt.Println("Typing via keyboard...")
}

func (a accountant) walking() {
	fmt.Println("The accountant comes to work on time")
}

func (a accountant) accrue() {
	fmt.Println("Accuals: mr.Presley + $1000, mr.Johnson + $2000")
}

func (a accountant) calculate() {
	fmt.Println("Current assets devide by current liabilities...")
}

func main() {
	var buh finansist

	buh = &accountant{experience: 15, age: 39}
	buh.typing()
	buh.walking()
	buh.accrue()
	buh.calculate()
}
