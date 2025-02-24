package main

import (
  "fmt"
  "modvar"
)

func main() {
  subscriber := modvar.Subscriber{Name: "Omar Khayyam"}
  subscriber.Street = "Sedova St" // .Address. omitted
  subscriber.City = "Kazan"       // .Address. omitted
  subscriber.State = "Tatarstan"  // .Address. omitted
  subscriber.PostalCode = "420133"// .Address. omitted

  fmt.Println("Street: ",  subscriber.Street)
  fmt.Println("City: ", subscriber.City)
  fmt.Println("State: ", subscriber.State)
  fmt.Println("PostalCode: ", subscriber.PostalCode)

  employee := modvar.Employee{Name: "Joe Satriani"}
  employee.Street = "456 Elm St"
  employee.City = "Portland"
  employee.State = "OR"
  employee.PostalCode = "97222"

  fmt.Println("Street: ", employee.Street)
  fmt.Println("City: ", employee.City)
  fmt.Println("State: ", employee.State)
  fmt.Println("PostalCode: ", employee.PostalCode)
}
