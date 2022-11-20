package main

import "fmt"

type IPAddr [4]byte

func (m map[string]int) string {
  
  }
// TODO: Add a "String() string" method

func main() {
  hosts := map[string]IPAddr{
    "loopback": {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }

  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}


// https://go.dev/tour/methods/18
