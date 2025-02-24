// https://networkbit.ch/golang-dns-lookup/
package main

import (
  "fmt"
  "net"
)

func main() {
  var hostname string
  fmt.Printf("Enter hostname: ")
  fmt.Scanf("%s\n", &hostname)

  printLookupIP(hostname)
}

func printLookupIP(hostname string) {
  ips, err := net.LookupIP(hostname)
  if err != nil {
    fmt.Println("LookupIP: ", err.Error())
  }

  if len(ips) == 0 {
    fmt.Println("no records")
  }

  for _, ip := range ips {
    fmt.Printf("%s\n", ip.String())
  }
  fmt.Println()
}
