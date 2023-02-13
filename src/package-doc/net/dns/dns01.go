// https://networkbit.ch/golang-dns-lookup/
package main

import (
  "fmt"
  "net"
)

func main() {
  var ip string
  fmt.Printf("Enter IP: ")
  fmt.Scanf("%s\n", &ip)
  printLookupAddr(ip)
}


func printLookupAddr(ip string) {
  names, err := net.LookupAddr(ip)
  if err != nil {
    fmt.Println("reverse lookup: ", err.Error())
  }

  if len(names) == 0 {
    fmt.Println("no records")
  }

  for _, name := range names {
    fmt.Printf("%s\n", name)
  }
  fmt.Println()
}


// net.LookupAddr() performs a reverse lookup for the given address, returning
// a list of names mapping to that address.
