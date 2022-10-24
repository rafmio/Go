package main

import (
    "fmt"
    "geo"
)

func main() {
    location := geo.Coordinates{Latitude: 37.42, Longitude: -122.0}
    fmt.Println("Latitude: ", location.Latitude)
    fmt.Println("Longitude: ", location.Longitude)

    fmt.Println()

    location2 := geo.Landmark{}
    location2.Name = "The Googleplex"
    location2.Latitude = 37.42
    location2.Longitude = -122.08
    fmt.Println(location2)
    fmt.Println(location2.Name)
    fmt.Println(location2.Latitude)
    fmt.Println(location2.Longitude)
}
