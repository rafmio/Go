// Empty struct
// https://dave.cheney.net/2014/03/25/the-empty-struct
package main

import "fmt"

type S struct{}

func (s *S) addr() {
  fmt.Printf("%p\n", s)
}

func main() {
  var a, b S
  a.addr()
  b.addr()

}
