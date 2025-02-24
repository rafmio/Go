package main

import "fmt"

type employee interface {
	typing()
	walking()
}

type finansist interface {
	employee
	accure()
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
	fmt.Println("Tne accountant comes to work on time")
}

func (a accountant) accure() {
	fmt.Println("Accruals: mr.Presley: + $1000, mr.Johnson + $2000")
}

func (a accountant) calculate() {
	fmt.Println("Current assets devide by current liabilities...")
}

func main() {
	var buh finansist

	buh = &accountant{experience: 15, age: 39}
	buh.typing()
	buh.walking()
	buh.accure()
	buh.calculate()
}
