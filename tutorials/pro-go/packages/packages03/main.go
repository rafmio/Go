package main

import (
	"fmt"
	_ "package03/data"
	. "package03/fmt"
	"package03/store"
	"package03/store/cart"
	"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279.00)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))

	color.Green("Name:" + cart.CustomerName)
	color.Cyan("Total:" + ToCurrency(cart.GetTotal()))

	color.Blue("Hello, I'm the the Blue!")
	color.HiBlack("HiBlack")
}
