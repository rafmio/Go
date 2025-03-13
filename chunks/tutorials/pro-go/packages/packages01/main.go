package main

import (
    "fmt"
    "packages01/store"
    currencyFmt "packages01/fmt" // Alias
//     . "package01/fmt" - dot import 
)

func main() {
    product := store.Product {
        Name: "Kayak",
        Category: "Watersports",
    }
    
    fmt.Println("Name:", product.Name)
    fmt.Println("Category:", product.Category)
    fmt.Println()
    
    fmt.Println("Price:",currencyFmt.ToCurrency(product.Price()))
//     Dot import:
//     fmt.Println("Price:", ToCurrency(product.Price()))
    fmt.Println()
// -----------------------------------------------------
    
    product2 := store.NewProduct("Submarine", "Watersports", 90000)
    
    fmt.Println("Name:", product2.Name)
    fmt.Println("Category:", product2.Category)
    fmt.Println("Price:", product.Price())
}

// "package01" - module
// "/" - separator
// "store" - package
// Exported types and fields should be started with 
// uppercase letter

// currencyFmt "packages/ftm"
// "currencyFmt" - alias
