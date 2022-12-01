package main

import "strconv"

type Product struct {
	Name     string
	Category string
	Price    float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279.00},
	{"Lifejacket", "Watersports", 49.50},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flag", "Soccer", 34.95},
	{"Stadium", "Soccer", 79_500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chear", "Chess", 75},
	{"Bling-bling King", "Chess", 1_200},
	{"Submarine", "Warship Games", 1000_000},
	{"Mine", "Warship Games", 1_000},
	{"Torpedo", "Warship Games", 1_700},
}

type ProductGroup []*Product
type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$ " + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
