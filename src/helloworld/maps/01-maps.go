// https://go.dev/tour/moretypes/19
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39968,
	}

	fmt.Println(m["Bell Labs"])

	var mm map[int]string
	fmt.Println("mm:", mm)
}
