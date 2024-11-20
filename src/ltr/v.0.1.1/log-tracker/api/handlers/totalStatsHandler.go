package handlers

import (
	"api/dbops"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type totalStatsParams struct {
	statIndicatorsNames   map[string]string    // map[statistic_name]"Statistic Print Name"
	statIndicatorsQueries map[string]string    // map[statistic_name]"SQL Query"
	statIndicatorsRows    map[string]*sql.Rows // the result of the SQL-query
}

type serverStat struct {
	serverName string // server name: 'cute_ganymede', 'black_oxygenium', etc
	totalStatsParams
	*dbops.DBConfig
	tsr            *totalStatsResult
	mapIpCountries map[string]string // map[ip_string]"Country Name"
	err            error
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

	dbConfigFilePath := "config/db-config.json"

	// new serverStatList []*serverStat
	serverStatList, err := newServerStatList(dbConfigFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, srv := range serverStatList {
		srv.totalStatsParams = *newTotalStatsParams() // filling totalStatsParams

		// set DB configs
		srv.DBConfig, err = dbops.NewDBConfig(dbConfigFilePath, srv.serverName)
		if err != nil {
			srv.err = err
		}
		srv.DBConfig.SetDSN()
		err = srv.DBConfig.EstablishDbConnection()
		if err != nil {
			http.Error(w, "Error: establishing connection to DB", http.StatusInternalServerError)
		}
		defer srv.DB.Close()

		srv.statIndicatorsRows = make(map[string]*sql.Rows, len(srv.statIndicatorsQueries))

		for qrNm, sqlQr := range srv.statIndicatorsQueries {
			resultOfSqlQr, err := srv.DB.Query(sqlQr)
			if err != nil {
				http.Error(w, "Error making SQL query", http.StatusInternalServerError)
			}
			srv.statIndicatorsRows[qrNm] = resultOfSqlQr
		}
	}

	for _, srv := range serverStatList {
		srv.tsr, err = extractStatIndicators(srv.statIndicatorsRows)
		if err != nil {
			http.Error(w, "Error extracting statistics", http.StatusInternalServerError)
		}
	}

	for _, srv := range serverStatList {
		fmt.Fprintf(w, "<h1>Server: %s ", srv.serverName)
		fmt.Fprintf(w, "Host: %s</h1>", srv.Host)
		fmt.Fprintf(w, "<p>DSN: %s</p>", srv.Dsn)
		fmt.Fprintf(w, "<h3>Total Stats:</h3>")
		fmt.Fprintf(w, "<p>Total records: %d</p>", srv.tsr.totalRecords)
		fmt.Fprintf(w, "<p>Unique IP Count: %d</p>", srv.tsr.uniqueIPCount)
		fmt.Fprintf(w, "<p>Records per day: %.2f</p>", srv.tsr.recordsPerDay)
		fmt.Fprintf(w, "<h4>Top 10 IPs:</h4>")
		for ip, num := range srv.tsr.topTenIPs {
			fmt.Fprintf(w, "<p>%s : %d</p>", ip, num)
		}
		fmt.Fprintf(w, "<h4>Top 10 Department ports:</h4>")
		for dpt, num := range srv.tsr.topTenDpt {
			fmt.Fprintf(w, "<p>%s : %d</p>", dpt, num)
		}

	}
}

func newServerStatList(dbConfigFilePath string) ([]*serverStat, error) {
	// open config file
	serverStatList := make([]*serverStat, 0)

	file, err := os.ReadFile(dbConfigFilePath)
	if err != nil {
		log.Println("opening config file:", err)
		return nil, err
	}

	// unmarshalling JSON data to struct
	dbConfigs := make(map[string]dbops.DBConfig) // variable for storing unmarshalled data
	err = json.Unmarshal(file, &dbConfigs)
	if err != nil {
		log.Println("Unmarshalling JSON:", err)
		return nil, err
	}

	for srvName := range dbConfigs { // see 'Simplify range: https://pkg.go.dev/golang.org/x/tools/gopls/internal/analysis/simplifyrange'
		if strings.Contains(srvName, "test") {
			continue
		} else {
			newServerStat := new(serverStat)
			newServerStat.serverName = srvName
			// append newServerStat to serverStatList
			serverStatList = append(serverStatList, newServerStat)
		}
	}

	return serverStatList, nil
}

func newTotalStatsParams() *totalStatsParams {
	tsp := new(totalStatsParams)
	tsp.statIndicatorsNames = make(map[string]string)
	tsp.statIndicatorsQueries = make(map[string]string)

	// map of stat indicator names
	tsp.statIndicatorsNames["total_records"] = "Total Number of Records"
	tsp.statIndicatorsNames["unique_ip_count"] = "Unique IP Count"
	tsp.statIndicatorsNames["records_per_day"] = "Records Per Day"
	tsp.statIndicatorsNames["top_10_ips"] = "Top 10 Most Frequent IP Addresses"
	tsp.statIndicatorsNames["top_10_dpt"] = "Top 10 Most Frequent Destination Ports"

	// set SQL queries for statistics
	tsp.statIndicatorsQueries["total_records"] = `SELECT COUNT(*) FROM lg_tab`
	tsp.statIndicatorsQueries["unique_ip_count"] = `SELECT COUNT(DISTINCT srcip) FROM lg_tab`

	// returned value: average records per day
	tsp.statIndicatorsQueries["records_per_day"] = `
	   SELECT AVG(daily_count) AS average_records_per_day
		FROM (
		    SELECT COUNT(*) AS daily_count
    		FROM lg_tab
    		GROUP BY DATE(tmstmp)
			) AS subquery;
    `

	// returned value: ip:count
	tsp.statIndicatorsQueries["top_10_ips"] = `
        SELECT srcip, COUNT(*) AS ip_count
        FROM lg_tab
        GROUP BY srcip
        ORDER BY ip_count DESC
        LIMIT 10;
    `

	// returned value: destination_port:count
	tsp.statIndicatorsQueries["top_10_dpt"] = `
	SELECT dpt, COUNT(*) AS count
	FROM lg_tab
	GROUP BY dpt
	ORDER BY count DESC
	LIMIT 10;
	`

	return tsp
}
