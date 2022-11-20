package main
import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")

  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X")
  case "linux":
    fmt.Printf("%s\n", os)
  default:
    fmt.Printf("%s\n", os)
  }

  var name string = "Mio"
  switch name {
  case "Viktor":
    fmt.Println("Victor")
  case "Nio":
    fmt.Println("It seems to be Mio may be?")
  case "Mio":
    fmt.Println("Indeed", name)

  }
}


// Switch
// https://go.dev/tour/flowcontrol/9
