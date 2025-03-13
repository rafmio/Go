package main

import (
    "fmt"
    "packages02/store"
    . "packages02/fmt"
    "packages02/store/cart"
)

func main() {
    product := store.NewProduct("Kayak", "Watersports", 279)
    
    cart := cart.Cart {
        CustomerName: "Alice", 
        Products: []store.Product{ *product },
    }
    
    fmt.Println("Name:", cart.CustomerName)
    fmt.Println("Total:", ToCurrency(cart.GetTotal()))
}
