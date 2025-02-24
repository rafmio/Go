// https://www.geeksforgeeks.org/time-newtimer-function-in-golang-with-examples/
package main

import (
  "fmt"
  "time"
)

func main() {
  newtimer := time.NewTimer(2 * time.Second)
  cNT := <- newtimer.C
  fmt.Println("Timer is inactivated")
  fmt.Println(cNT)

  for i := 0; i < 3; i++ {
    fmt.Println(i, "Wating for 'Sleep()'...")
    time.Sleep(time.Second * 1)
  }

}
