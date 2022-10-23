package main

import (
    "fmt"
    "modvar"
)

func main() {
    address := modvar.Address{Street: "123 Oak St", City: "Omaha", 
        State: "NE", PostalCode: "68111"}
    subscriber := modvar.Subscriber{Name: "Aman Singh"}
    subscriber.HomeAddress = address
    fmt.Println(subscriber)
    fmt.Println(subscriber.HomeAddress)
}
