package models

type Icecream struct {
	Flavour   string
	Scoops    int
	BasePrice float32
}

func (i Icecream) Price() float32 {
	return i.BasePrice * float32(i.Scoops)
}
