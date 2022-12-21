// Using struct tags to control decoding
package main

import "encoding/json"

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:"offer,string"`
}

func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}

// The tag applied to the Discount field tells the Decoder that the value for
// this field should be obtained from the JSON key named 'offer' and that the
// value will be parsed from a string, instead of the JSON number
