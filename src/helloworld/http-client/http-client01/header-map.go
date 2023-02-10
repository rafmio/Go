package main

import (
  "fmt"
)

func main() {
  headers := map[string][]string {
    "Accept": {"application/json", "text/javascript"},
    "Accept-Encoding": {"gzip", "deflate", "br"},
    "Content-Type": {"application/x-www-form-urlencoded; charset=UTF-8"} // ?? точки с запятой - отдельные элементы слайса?
    
  }

  fmt.Println(headers)
}
