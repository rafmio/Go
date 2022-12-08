// https://go.dev/blog/strings
package main

import "fmt"

const placeOfInterest = `⌘`


func main() {
  fmt.Printf("plaing string: ")
  fmt.Printf("%s", placeOfInterest)
  fmt.Printf("\n")

  fmt.Printf("quoted string: ")
  fmt.Printf("%+q", placeOfInterest)
  fmt.Printf("\n")

  fmt.Printf("hex bytes: ")
  for i := 0; i < len(placeOfInterest); i++ {
    fmt.Printf("%x ", placeOfInterest[i])
  }
  fmt.Printf("\n")
}
