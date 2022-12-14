package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "net/http"
  "html/template"
)
//----------------------------------------------------------------------------
type Guestbook struct {
  SignatureCount int
  Signatures []string
}
//----------------------------------------------------------------------------
func getStrings(fileName string) []string {
  var lines []string
  file, err := os.Open(fileName)
  if os.IsNotExist(err) {
    return nil
  }
  check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  check(scanner.Err())
  return lines
}
//----------------------------------------------------------------------------
func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
//----------------------------------------------------------------------------
func viewHandler(writer http.ResponseWriter, request *http.Request) {
  signatures := getStrings("signatures.txt")
  fmt.Printf("%#v\n", signatures)
  html, err := template.ParseFiles("view.html")
  check(err)
  guestbook := Guestbook{
    SignatureCount: len(signatures),
    Signatures: signatures,
  }
  err = html.Execute(writer, guestbook)
  check(err)
}
//----------------------------------------------------------------------------
func newHandler(writer http.ResponseWriter, request *http.Request) {
  html, err := template.ParseFiles("new.html")
  check(err)
  err = html.Execute(writer, nil)
  check(err)
}
//----------------------------------------------------------------------------
func main() {
  http.HandleFunc("/guestbook", viewHandler)
  http.HandleFunc("/guestbook/new", newHandler)
  err := http.ListenAndServe("localhost:8080", nil)
  log.Fatal(err)
}
