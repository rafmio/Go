package main

import (
  "fmt"
  "time"
)

// func receiveDispatches(channel <- chan DispatchNotification) {
//   for details := range channel {
//     fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
//     "x", details.Product.Name)
//   }
//   fmt.Println("Channel has been closed")
// }
func enumerateProducts(channel chan <- *Product) {
  for _, p := range ProductList[:3] {
    channel <- p
    time.Sleep(time.Millisecond * 800)
  }
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
  // receiveDispatches((<- chan DispatchNotification)(dispatchChannel))
  for {
    select {
    case details, ok := <- dispatchChannel:
      if ok {
        fmt.Println("Dispatch to", details.Customer, ":",
        details.Quantity, "x", details.Product.Name)
      } else {
        fmt.Println("Channel has been closed")
        goto alldone
      }
    default:
      fmt.Println("-- No message ready to be received")
      time.Sleep(time.Millisecond * 500)
    }
  }
  alldone: fmt.Println("All values received")
}
