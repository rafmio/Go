// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	indianTaxI := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}

	singaporeTaxI := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	usaTaxI := &usaTax{
		taxPercentage: 40,
		income:        3000,
	}

	taxSystems := []taxSystem{indianTaxI, singaporeTaxI, usaTaxI}
	totalTax := calculateTotalTax(taxSystems)
	fmt.Println("Total Tax is: ", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax()
	}
	return totalTax
}
