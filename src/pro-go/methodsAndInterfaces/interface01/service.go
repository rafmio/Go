package main

type Service struct {
  description string
  durationMonths int
  monthlyFee float64
}

func (s Service) getName() string {
  return s.description
}

func (s Service) getCost(recur bool) float64 {
  if (recur) {
    return s.monthlyFee * float64(s.durationMonths)
  }
  return s.monthlyFee
}


// To implement an interface, all the methods specified by the interface
// must be defined for a struct type
// Methods must have the same name, paremeter types, and result types
