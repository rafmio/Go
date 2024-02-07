package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	formUrl := &url.URL{
		Scheme: "http",
		Host:   "himpostavka.ru",
	}

	link := formUrl.String()
	fmt.Println(link)

	rsp, _ := http.Get(link)
	fmt.Println(rsp.Status)

	for _, val := range rsp.Header {
		fmt.Println(val)
	}
}
