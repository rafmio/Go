package main

import "errors"

type ChannelMessage struct {
	Category string
	Total    float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice(category string) (total float64,
err error) {
	productCount := 0
	
}
