package main

import "fmt"

func ToGallons(l Liters) Gallons {
  return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
  return Liters(g * 3.785)
}

func main() {
  type Liters
  carFuel := Gallons(1.2)
  busFuel := Liters(4.5)
}
