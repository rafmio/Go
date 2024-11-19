package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type countryName struct {
	ip      string `json:"ip"`
	country string `json:"country"`
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
		resp, err := http.Get(fmt.Sprintf(url, ip))
		if err != nil {
			fmt.Printf("Error fetching IP information for %s: %v\n", ip, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for %s: %v\n", ip, err)
			continue
		}

		// fmt.Printf(string(body))

		cn := extractCountry(body)
		countries[cn.ip] = cn.country
	}

	fmt.Println("len(countries):", len(countries))
	for ip, country := range countries {
		fmt.Printf("IP: %s, Country: %s\n", ip, country)

	}
}

func extractCountry(body []byte) countryName {
	fmt.Println(string(body))
	cn := new(countryName)
	err := json.Unmarshal(body, cn)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return *cn
	}
	return *cn

}
