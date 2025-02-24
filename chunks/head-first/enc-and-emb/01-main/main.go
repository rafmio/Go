package main

import (
    "fmt"
    "geo"
)

func main() {
    coordinates := geo.Coordinates{}
    coordinates.SetLatitude(37.42)
    coordinates.SetLongitude(-122.48)
    fmt.Println(coordinates)
}
