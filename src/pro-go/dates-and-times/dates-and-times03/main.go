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
  layout := "2006-Jan-02"
  dates := []string {
    "1995-Jun-09",
    "2015-Jun-02",
  }

  for _, d := range dates {
    time, err := time.Parse(layout,d)
    if (err == nil) {
      PrintTime("Parsed", &time)
    } else {
      Printfln("Error: %s", err.Error())
    }
  }
  fmt.Println()

  dates2 := []string {
    "09 Jun 95 00:00 GMT",
    "02 Jun 15 00:00 GMT",
  }
  for _, d := range dates2 {
    time, err := time.Parse(time.RFC822, d)
    if (err == nil) {
      PrintTime("Parsed", &time)
    } else {
      Printfln("Error %s", err.Error())
    }
  }

}
