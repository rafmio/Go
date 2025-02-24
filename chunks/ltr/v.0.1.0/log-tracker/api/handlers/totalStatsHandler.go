package handlers

import "net/http"

type totalStatsParams struct {
	statIndicatorsNames   map[string]string // map[statistic_name]"Statistic Print Name"
	statIndicatorsQueries map[string]string // map[statistic_name]"SQL Query"
}

func TotalStatsHandler(w http.ResponseWriter, r *http.Request) {
	// processURL makes:
	// 	1. checking whether the HTTP request method is a GET method
	// 	2. setting headers
	// 	3. parsing request's parameters and populate r.Form
	err := processURL(w, r)
	if err != nil {
		if err == http.ErrNotSupported {
			http.Error(w, "only GET requests are supported", http.StatusMethodNotAllowed)
		}
		if err == ErrParsingForm {
			http.Error(w, "error parsing form", http.StatusBadRequest)
		}
	}
}

// QUERIES FOR STATISTICS
// g.queryStatIndicators["totalNumberEntries"] = `SELECT COUNT(*) FROM lg_tab`
// g.queryStatIndicators["uniqueIpCount"] = `SELECT COUNT(DISTINCT srcip) FROM lg_tab`
// g.queryStatIndicators["entriesPerDay"] = `
// 	SELECT AVG(daily_count) AS average_records_per_day
// 	FROM (
// 		SELECT COUNT(*) AS daily_count
// 		FROM lg_tab
// 		GROUP BY DATE(tmstmp)
// 		`
