package main

import (
    "fmt"
    "geo"
)

func main() {
    location := geo.Coordinates{Latitude: 37.42, Longitude: -122.0}
    fmt.Println("Latitude: ", location.Latitude)
    fmt.Println("Longitude: ", location.Longitude)
}
