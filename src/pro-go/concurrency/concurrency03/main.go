package main

import (
  "fmt"
)

func receiveDispatches(channel <- chan DispatchNotification) {
  for details := range channel {
    fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
    "x", details.Product.Name)
  }
  fmt.Println("Channel has been closed")
}

func main() {
  // --------------------------------------------------------
  // dispatchChannel := make(chan DispatchNotification, 100)
  //
  // go DispatchOrders(dispatchChannel)
  //
  // for details := range dispatchChannel {
  //   fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
  //   "x", details.Product.Name, "Catgory:", details.Product.Category)
  // }
  // fmt.Println("Channel has been closed")
  // --------------------------------------------------------
  // dispatchChannel := make(chan DispatchNotification, 100)
  //
  // var sendOnlyChannel chan <- DispatchNotification = dispatchChannel
  // var receiveOnlyChannel <- chan DispatchNotification = dispatchChannel
  //
  // go DispatchOrders(sendOnlyChannel)
  // receiveDispatches(receiveOnlyChannel)

  // Explict conversion: -------------------------------------------
  dispatchChannel := make(chan DispatchNotification, 100)
  go DispatchOrders(chan <- DispatchNotification(dispatchChannel))
  receiveDispatches((<- chan DispatchNotification)(dispatchChannel))

  
}
