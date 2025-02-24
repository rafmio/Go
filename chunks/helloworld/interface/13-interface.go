// https://golangbyexample.com/interface-in-golang/
package main

import (
	"fmt"
)

type taxSystem interface {
	calculateTax() int
}

type tax struct {
	location      string
	taxPercentage int
	income        int
}

func main() {
	indianTax := &tax{
		location:      "India",
		taxPercentage: 30,
		income:        1000,
	}

	singaporeTax := &tax{
		location:      "Singapore",
		taxPercentage: 10,
		income:        2000,
	}

	usTax := &tax{
		location:      "US",
		taxPercentage: 40,
		income:        3000,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax, usTax}
	totalTax := calculateTotalTax(taxSystems)
	fmt.Println("Total tax is: ", totalTax)
}

func (i *tax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax()
	}

	return totalTax
}
