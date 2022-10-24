package main

import (
    "errors"
    "fmt"
    "log"
)

type Date struct {
    Year int
    Month int
    Day int
}

func (d *Date) SetYear(year int) error {
    if year < 1970 {
        return errors.New("invalid year")
    }
    d.Year = year
    return nil
}

func (d *Date) SetMonth(month int) error {
    if month < 1 || month > 12 {
        return errors.New("invalid month")
    }
    d.Month = month
    return nil
}

func (d *Date) SetDay(day int) error {
    if day < 1 || day > 31 {
        return errors.New("invalid day")
    }
    d.Day = day 
    return nil
}

func main() {
    date := Date{}
    var year int
    fmt.Scan(&year)
    
    err := date.SetYear(year)
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("year: ", date.Year)
    }
    
    var month int
    fmt.Scan(&month)
    err = date.SetMonth(month)
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("month: ", date.Month)
    }
    
    var day int
    fmt.Scan(&day)
    err = date.SetDay(day)
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("day: ", date.Day)
    }
}
