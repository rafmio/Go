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
    if details, open := <-dispatchChannel; open {
      fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
      "x", details.Product.Name, "Category:", details.Product.Category)
    } else {
      fmt.Println("Channel has been closed")
      break
    }
  }
}
