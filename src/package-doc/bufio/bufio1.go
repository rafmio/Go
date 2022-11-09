package main

import (
  "bufio"
  "fmt"
  "strconv"
  "strings"
)

func main() {
  // An artificial inpurt source
  const input = "1234 5678 123456788909"
  scanner := bufio.NewScanner(strings.NewReader(input))

  // Create a custom split function by wrapping the existing ScanWords
  // function
  split := func(data []byte, atEOF bool)(advance int, token []byte, err error) {
    advance, token, err = bufio.ScanWords(data, atEOF)
    if err == nil && token != nil {
      _, err = strconv.ParseInt(string(token), 10, 32)
    }
    return
  }

  // Set the split function for the scanning operation
  scanner.Split(split)

  // Validate the input
  for scanner.Scan() {
    fmt.Printf("%s\n", scanner.Text())
  }
}
