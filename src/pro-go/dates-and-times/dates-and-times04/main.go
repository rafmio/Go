package main

import (
  "fmt"
  "time"
)

func PrintTime(label string, t *time.Time) {
  fmt.Println(label, t.Format(time.RFC822Z))
}

func main() {
  layout := "02 Jan 06 15:04"
  date := "09 Jun 95 19:30"

  london, lonner := time.LoadLocation("Europe/London")
  newyork, nycerr := time.LoadLocation("America/New_York")

  if (lonner == nil && nycerr == nil) {
    nolocation, _ := time.Parse(layout, date)
    londonTime, _ := time.ParseInLocation(layout, date, london)
    newyorkTime, _ := time.ParseInLocation(layout, date, newyork)

    PrintTime("No location:", &nolocation)
    PrintTime("London:", &londonTime)
    PrintTime("New York:", &newyorkTime)
  } else {
    fmt.Println(lonner.Error(), nycerr.Error())
  }

}
