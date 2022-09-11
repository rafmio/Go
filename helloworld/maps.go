package main
import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.86433, -74.3996,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m)
}


// Maps
// A map maps keys to values
// The zero value of a map is nil. A nil map has no keys, nor 
// can keys be added
// The 'make' function returns a map of the given type, 
// initialize and ready for use

// Maps in Golang
// https://golangdocs.com/maps-in-golang
