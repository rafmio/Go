package main

import (
  "fmt"
  "log"
  "mypackage"
)

func main() {
  event := mypackage.Event{}
  err := event.SetTitle("Mom's birthday")
  if err != nil {
    log.Fatal(err)
  }

  err = event.SetYear(2022)
  if err != nil {
    log.Fatal(err)
  }

  err = event.SetMonth(10)
  if err != nil {
    log.Fatal(err)
  }

  err = event.SetDay(26)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(event.Title())
  fmt.Println(event.Year())
  fmt.Println(event.Month())
  fmt.Println(event.Day())
}
