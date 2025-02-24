package main

import (
  "fmt"
  "time"
)

func PrintTime(label string, t *time.Time) {
  // layout := "Minute 5 Hour: 3 Day: 02 Month: Jan Year: 2006"
  // fmt.Println(label, t.Format(layout))
  fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
  current := time.Now()
  specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
  unix := time.Unix(1233228090, 0)

  PrintTime("Current ", &current)
  PrintTime("Specific", &specific)
  PrintTime("UNIX    ", &unix)
  fmt.Println()

  fmt.Println(current)
  fmt.Println(specific)
  fmt.Println(unix)
  fmt.Println(current.Year(), current.Month(), current.Day())
  fmt.Printf("Type of current:  %T\n", current)
  fmt.Printf("Type of specific: %T\n", specific)
  fmt.Printf("Type of unix:     %T\n", unix)
  fmt.Println()


}
