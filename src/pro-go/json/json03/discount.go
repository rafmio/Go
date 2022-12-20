package main

type DiscountedProduct struct {
	*Product `json:"product"`
	Discount float64
}

// `json:"product"` - tag for JSON encoding
