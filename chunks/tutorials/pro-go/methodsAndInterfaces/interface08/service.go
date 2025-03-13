package main

import "fmt"

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

func (s Service) getCostAndPrice(recur bool) {
	if recur {
		fmt.Println("Service:", s.description,
			"Price:", s.monthlyFee * float64(s.durationMonths))
	}
	fmt.Println("Service:", s.name, "Price:", s.monthlyFee)
}
