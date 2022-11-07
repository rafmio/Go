package prose

import (
  "testing"
  "fmt"         // fmt package is for Sprintf()
)

func TestTwoElements(t *testing.T) {
  list := []string{"apple", "orange"}
  want := "apple and orange"
  got := JoinWithCommas(list)
  if got != want {
    t.Error(list, got, want)
  }
}

func TestThreeElements(t *testing.T) {
  list := []string{"apple", "orange", "pear"}
  want := "apple, orange, and pear"
  got := JoinWithCommas(list)
  if got != want {
    t.Error(list, got, want)
  }
}

func errorString(list []string, got string, want string) string {
  return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
