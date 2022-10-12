package main

import (
  "fmt"
  "dates"
)

func main() {
  days := 3
  fmt.Println("Your appointment in in", days, "days")
  fmt.Println("with a follow-up in", days + dates.DaysInWeek, "days")
}
