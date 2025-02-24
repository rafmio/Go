package main
import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }

  return lim
}

func chckval(xx int) string {
  if a := xx; a < 150 {
    var msg string = "Checked variable still less than 150"
    return msg
  }
  return "Checked variable is greater than 150"
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )

  var chck1 int = 200
  var reschck1 string = chckval(chck1)
  fmt.Println(reschck1)

}


// If with a short statement
// https://go.dev/tour/flowcontrol/6
