// https://www.geeksforgeeks.org/time-newtimer-function-in-golang-with-examples/
package main

import (
  "fmt"
  "time"
)

func main() {
  newtimer := time.NewTimer(time.Second)

  go func() {             // Notifying channel under go function
    <- newtimer.C
    fmt.Println("timer inactivated") // Printed when timer is fired
  }()

  stopTimer := newtimer.Stop()

  if stopTimer {
    fmt.Println("The timer is stopped")
  }
  time.Sleep(time.Second * 2)
}
