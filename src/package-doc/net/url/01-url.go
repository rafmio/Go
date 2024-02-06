package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	u := `himpostavka.ru`

	// if !strings.Contains(u, `://`) {
	// 	u = `http://` + u
	// }

	if !strings.HasPrefix(strings.ToLower(u), `http://`) {
		u = `http://` + u
	}

	rsp, err := http.Get(u)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("The response is not empty")
		fmt.Println(rsp)
	}

	defer rsp.Body.Close()

}
