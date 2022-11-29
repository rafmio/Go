package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("main function started")
  CalcStoreTotal(Products)
  time.Sleep(time.Second * 4)
  fmt.Println("main function complete")
}
