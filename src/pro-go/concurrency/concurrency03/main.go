package main

import (
  "fmt"
)

func main() {
  dispatchChannel := make(chan DispatchNotification, 100)

  go DispatchOrders(dispatchChannel)
  
  for details := range dispatchChannel {
    fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
    "x", details.Product.Name, "Catgory:", details.Product.Category)
  }
  fmt.Println("Channel has been closed")
}
