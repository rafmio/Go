package main

import (
    "fmt"
    "log"
    "mypackage"
)

func main() {
    event := mypackage.Event{}
    err := event.SetYear(2019)
    if err != nil {
      log.Fatal(err)
    }

    err = event.SetMonth(5)
    if err != nil {
      log.Fatal(err)
    }

    err = event.SetDay(27)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(event.Year())
    fmt.Println(event.Month())
    fmt.Println(event.Day())

    fmt.Println()
    event.Date.Hour = 12
    fmt.Println(event.Date.Hour)

    fmt.Println(event.Date.Year())
    fmt.Println(event.Date.Month())
    fmt.Println(event.Date.Day())

    event.Title = "Mom's bithday"
    fmt.Println()
    fmt.Println(event.Title)
}
