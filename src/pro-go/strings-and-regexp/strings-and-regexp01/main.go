package main

import (
  "fmt"
  "strings"
  "bytes"
)

func main() {
  product := "Kayak"
  fmt.Println("Contains:", strings.Contains(product, "yak"))
  fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
  fmt.Println("ContainsRune", strings.ContainsRune(product, 'K'))
  fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
  fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
  fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
  fmt.Println()

  price := "€100"
  fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
  fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
  []byte {226, 130}))
  fmt.Println()

  description := "A boat for sailing"
  fmt.Println("Original:", description)
  fmt.Println("Title:", strings.Title(description))
  fmt.Println("ToTitle:", strings.ToTitle(description))
  fmt.Println("ToLower:", strings.ToLower(description))
  fmt.Println("ToUpper:", strings.ToUpper(description))
  fmt.Println("ToLower again:", strings.ToLower(description))
  fmt.Println("Title again:", strings.Title(description))
  fmt.Println()

  specialChar := "\u01c9"
  fmt.Println("Original:", specialChar, []byte(specialChar))
  upperChar := strings.ToUpper(specialChar)
  fmt.Println("Upper:", upperChar, []byte(upperChar))
  titleChar := strings.ToTitle(specialChar)
  fmt.Println("Title:", titleChar, []byte(titleChar))
  fmt.Println()
}
