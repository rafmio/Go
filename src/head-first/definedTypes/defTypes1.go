package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
  var carFuel Gallons
  var busFuel Liters
  carFuel = Gallons(10.0)
  busFuel = Liters(240.0)
  fmt.Println(carFuel, busFuel)

  carFuel2 := Gallons(10.0)
  busFuel2 := Liters(240.0)
  fmt.Println(carFuel2, busFuel2)

  carFuel = Gallons(Liters(40.0) * 0.264)
  busFuel = Liters(Gallons(63.0) *  3.785)
  fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

  fmt.Printf("Type of carFuel: %T\nType of busFuel: %T\n", carFuel, busFuel)

  fmt.Println(Liters(1.2) + Liters(3.4))
  var ltrs Liters
  ltrs = busFuel * Liters(2)
  fmt.Println(ltrs)

  fmt.Println()
  fmt.Println(Gallons(5.5) - Gallons(2.2))
  fmt.Println(Liters(2.2) / Liters(1.1))
  fmt.Println(Gallons(1.2) == Gallons(1.2))
  fmt.Println(Liters(2.2) < Liters(3.4))
  fmt.Println(Liters(1.2) > Liters(3.4))
}
