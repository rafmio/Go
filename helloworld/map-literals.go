package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex {
  "Bell Labs" : Vertex {
    40.68433, -74.39967,
  },

  "Google": Vertex {
    37.42202, -122.08408,
  },


}

func main() {
  fmt.Println(m)

  var mm = map[int]string {
    0: "Kissy",
    1: "Missy",
  }

  fmt.Println("mm : ", mm)
}


// Map literals
// https://go.dev/tour/moretypes/20
