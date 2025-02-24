// https://networkbit.ch/golang-dns-lookup/
package main

import (
  "fmt"
  "net"
)

func main() {
  var hostame string
  fmt.Printf("Enter hostame: ")
  fmt.Scanf("%s\n", &hostame)
  printLookupTXT(hostame)
}

func printLookupTXT(hostame string) {
  txts, err := net.LookupTXT(hostame)
  if err != nil {
    fmt.Println("LookupTXT: ", err.Error())
  }
  if len(txts) == 0 {
    fmt.Println("no records")
  }

  for _, txt := range txts {
    fmt.Printf("%s\n", txt)
  }
}
