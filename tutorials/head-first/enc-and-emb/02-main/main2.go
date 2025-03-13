package main

import (
  "fmt"
  "calendar"
  "log"
)

func main() {
  // Declare date instance: --------------------------------------------------
  date := calendar.Date{}

  // Set date field via SetYear method ---------------------------------------
  err := date.SetYear(2022)
  if err != nil {
    log.Fatal(err)
  }

  // Set month field via SetMonth method -------------------------------------
  err = date.SetMonth(10)
  if err != nil {
    log.Fatal(err)
  }

  // Set day field via SetDay method -----------------------------------------
  err = date.SetDay(25)
  if err != nil {
    log.Fatal(err)
  }

  // Print date's fields -----------------------------------------------------
  fmt.Println(date.Year())
  fmt.Println(date.Month())
  fmt.Println(date.Day())
}
