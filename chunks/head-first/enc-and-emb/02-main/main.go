package main

import (
  "fmt"
  "calendar"
  "log"
)

func main() {
  date := calendar.Date{}

  err := date.SetYear(2022)
  if err != nil {
    log.Fatal(err)
  }
  err = date.SetMonth(10)
  if err != nil {
    log.Fatal(err)
  }
  err = date.SetDay(31)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(date)
  
}


// Export struct, export methods, but don't export struct's fields
// Поля структуры Date могут обновляться только set-методами, программа
// защищена от случайного ввода недействительных данных
// НО!
// При попытке чтения отдельных полей структуры окажется, что поля неэкспорти-
// руемы и прочесть их невозможно
