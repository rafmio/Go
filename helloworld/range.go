package main
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 512}
var msg = []string{"Huggy", "Wuggy", "Kissy", "Missy", "Toko", "Boko"}

func main() {
  for i, v := range pow {
    fmt.Printf("2 ** %d = %d\n", i, v)
  }

  for indx, inmsg := range msg {
    fmt.Printf("index: %d, message: %s\n", indx, inmsg)
  }

}


// Range
// https://go.dev/tour/moretypes/16
// Return value:
// First: index
// Second: copy of the element at that index
