// https://networkbit.ch/golang-dns-lookup/
package main

import (
  "fmt"
  "net"
)

func main() {
  var host string
  fmt.Printf("Enter host name: ")
  fmt.Scanf("%s\n", &host)
  printLookupHost(host)
}

func printLookupHost(host string) {
  ips, err := net.LookupHost(host)
  if err != nil {
    fmt.Println("LookupHost: ", err.Error())
  }
  for _, ip := range ips {
    fmt.Printf("%s\n", ip)
  }
  fmt.Println()
}
