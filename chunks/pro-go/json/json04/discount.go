package main

// type DiscountedProduct struct {
//   *Product `json:"product"`
//   Discount float64 `json:"-"`
// }

type DiscountedProduct struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:"-"`
}

// `json:"product"` - tag for JSON encoding

// To skip a nil field without changing the name or field promotion,
// specify the omitempty keyword without a name:
// type DiscountedProduct struct {
//   *Product `json:",omitempty"`
//   Discount float64 `json:"-"`
// }
