package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  client := &http.Client{Timeout: time.Second * 10}
  var url string = "hhttps://xn--90adear.xn--p1ai"
  reqest, err := http.NewRequest("POST", url, nil)

  reqest.Header["Accept"] = {"application/json", "text/javascript", "*/*; q=0.01"}

}
  // Request Headers:

  // Accept: application/json, text/javascript, */*; q=0.01
  // Accept-Encoding: gzip, deflate, br
  // Accept-Language: en,ru;q=0.9
  // Connection: keep-alive
  // Content-Length: 122
  // Content-Type: application/x-www-form-urlencoded; charset=UTF-8
  // Host: xn--b1afk4ade.xn--90adear.xn--p1ai
  // Origin: https://xn--90adear.xn--p1ai
  // Referer: https://xn--90adear.xn--p1ai/
  // sec-ch-ua: "Chromium";v="106", "Yandex";v="22", "Not;A=Brand";v="99"
  // sec-ch-ua-mobile: ?0
  // sec-ch-ua-platform: "Linux"
  // Sec-Fetch-Dest: empty
  // Sec-Fetch-Mode: cors
  // Sec-Fetch-Site: same-site
  // User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 YaBrowser/22.11.3.838 Yowser/2.5 Safari/537.36
}
