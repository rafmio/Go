package main

import (
    "fmt"
    "modvar"
)

func main() {
    subscriber := modvar.Subscriber{Name: "Charles Dickens"}
    subscriber.Address.Street = "123 Oak St"
    subscriber.Address.City = "Omaha"

    fmt.Println("Street: ", subscriber.Address.Street)
    fmt.Println("City: ", subscriber.Address.City)

    employee := modvar.Employee{Name: "Joy Carr"}
    employee.Address.State = "OR"
    employee.Address.PostalCode = "97222"

    fmt.Println("State: ", employee.Address.State)
    fmt.Println("PostalCode: ", employee.Address.PostalCode)
}

// Anonymous struct fields
