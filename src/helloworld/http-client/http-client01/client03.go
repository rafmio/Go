// https://networkbit.ch/golang-http-client/
package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "net/http"
  "time"
)

func main() {
  url := "https://httpbin.org/post"
  data := []byte(`{"Hello": "world"}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
  if err != nil {
    log.Fatal("Error reading request: ", err)
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Host", "rafmio.ru")
  req.Header.Set("User-Agent", "Mozilla 5.0")

  cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
  req.AddCookie(&cookie)

  client := &http.Client{Timeout: time.Second * 10}

  fmt.Println(req.Cookies())
  fmt.Println(req.Header)

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Error sending request or reading response: ", err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status: ", resp.Status)
  fmt.Println("response Headers: ", resp.Header)

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error reading body: ", err)
  }

  fmt.Printf("Body: \n %s\n End of Body\n", body)
}
