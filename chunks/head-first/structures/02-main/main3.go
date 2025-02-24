package main

import (
  "fmt"
  "modvar"
)

func main() {
  subscriber := modvar.Subscriber{Name: "Aman Shaman"}
  subscriber.HomeAddress.Street = "123 Oad St"
  subscriber.HomeAddress.City = "Omaha"
  subscriber.HomeAddress.State = "NE"
  subscriber.HomeAddress.PostalCode = "68111"

  fmt.Println("Subscriber Name: ", subscriber.Name)
  fmt.Println("Street: ", subscriber.HomeAddress.Street)
  fmt.Println("City: ", subscriber.HomeAddress.City)
  fmt.Println("State: ", subscriber.HomeAddress.State)
  fmt.Println("PostalCode: ", subscriber.HomeAddress.PostalCode)
}
