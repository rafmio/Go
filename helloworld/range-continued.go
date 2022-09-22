package main
import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i)     // == 2 ** i
  }

  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }

  fmt.Println("pow: ", pow)

  sls1 := make([]int, 0)

  fmt.Println("sls1: ", sls1)
  sls1 = append(pow)
  fmt.Println("sls1: ", sls1)
  fmt.Println("pow: ", pow)

  for _, vl := range sls1 {
    fmt.Println("sls1: ", vl)
  }

  for indx, _ := range sls1 {
    fmt.Println("sls1: ", indx)
  }
}



// Range
// https://go.dev/tour/moretypes/17
