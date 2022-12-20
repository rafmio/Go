package main

// Forcing Fields to be Encoded as Strings
type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"`
}
