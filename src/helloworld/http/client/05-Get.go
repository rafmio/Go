package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	IP      string `json:"ip"`
	Success bool   `json:"success"`
	// Type        string `json:"type"`
	Country string `json:"country"`
	// CountryCode string `json:"country_code"`
}

func main() {
	IPs := []string{
		"194.58.102.129",
		"147.45.226.19",
		"109.205.213.106",
		"79.124.49.130",
		"103.207.36.205",
		"91.191.209.46",
		"103.153.77.178",
		"109.205.213.222",
	}

	countries := make(map[string]string)

	url := "http://ipwho.is/%s"

	for _, ip := range IPs {
		req, err := http.Get(fmt.Sprintf(url, ip))
		if err != nil {
			fmt.Printf("Error fetching IP information for %s: %v\n", ip, err)
			continue
		}
		defer req.Body.Close()

		var response Response
		if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
			fmt.Println("Error decoding IP information: ", err)
			return
		}

		if !response.Success {
			fmt.Println("Error retrieving IP information for", ip)
		}

		countries[ip] = response.Country

	}
	fmt.Println("Countries for given IP addresses:")
	for ip, country := range countries {
		fmt.Printf("%s: %s\n", ip, country)
	}

}
