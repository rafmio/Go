// https://jonathanmh.com/p/go-import-file-from-sub-directory/
package main

import (
	"fmt"
	"icecream-shop/models"
)

func main() {
	myIcecream := models.Icecream{Flavour: "vanilla", Scoops: 3, BasePrice: 2.99}
	fmt.Println(myIcecream.Price())
}
