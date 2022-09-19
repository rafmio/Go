package main
import (
  "fmt"
  "time"
)

// for Example 1----------------------------
func display(str string) {
  for w := 0; w < 6; w++ {
    time.Sleep(1 * time.Second)
    fmt.Println(str)
  }
}

func main() {
  // Example 1 -------------------------------
  go display("Welcome")
  display("Come here to me")
}
