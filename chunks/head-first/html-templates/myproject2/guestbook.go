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

//---Get strings--------------------------------------------------------------
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

//---CHECK--------------------------------------------------------------------
func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
//---View Hanlder-------------------------------------------------------------
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
//---main---------------------------------------------------------------------
func main() {
  http.HandleFunc("/guestbook", viewHandler)
  err := http.ListenAndServe("localhost:8080", nil)
  log.Fatal(err)
}
