package prose

import (
  "strings"
  "fmt"
)

func JoinWithCommas(phrases []string) string {
  result := strings.Join(phrases[:len(phrases) - 1], ", ")
  fmt.Println("1: print inside func: ", result)
  result += " and "
  fmt.Println("2: print inside func: ", result)
  result += phrases[len(phrases) - 1]
  fmt.Println("3: print inside func: ", result)
  return result
}
