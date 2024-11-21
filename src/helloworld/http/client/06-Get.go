package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type totalStatsResult struct {
	TotalRecords   int               `json:"total_records"`
	UniqueIPCount  int               `json:"unique_ip_count"`
	RecordsPerDay  float64           `json:"records_per_day"`
	TopTenIPs      map[string]int    `json:"top_10_ips"`
	TopTenDpt      map[string]int    `json:"top_10_dpt"`
	MapIpCountries map[string]string `json:"map_ip_countries"`
}

func main() {
	url := "http://localhost:8082/logtracker/statistics/total_stats"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	resultSls := make([]*totalStatsResult, 2)

	if err := json.NewDecoder(resp.Body).Decode(&resultSls); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for i, val := range resultSls {
		fmt.Println(i, ":", val)
	}
}
