package main

import (
  "fmt"
  // "time"
)

func main() {
  fmt.Println("main function started")

  // fmt.Println("Print 'Products':")
  // fmt.Println(Products)
  // fmt.Println()

  CalcStoreTotal(Products)
  // time.Sleep(time.Second * 4)
  fmt.Println("main function complete")

  fmt.Println("------------------------------------------------")
  fmt.Println()

  dispatchChannel := make(chan DispatchNotification, 100)
  go DispatchOrders(dispatchChannel)
  for {
    details := <- dispatchChannel
    fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
      "x", details.Product.Name)
  }
}
