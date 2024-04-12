package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//   layout := "2006-01-02"
	timeStc := "Apr 11 23:03:43"

	months := make(map[string]string, 12)
	months["Jan"] = "01"
	months["Feb"] = "02"
	months["Mar"] = "03"
	months["Apr"] = "04"
	months["May"] = "05"
	months["Jun"] = "06"
	months["Jul"] = "07"
	months["Aug"] = "08"
	months["Sep"] = "09"
	months["Oct"] = "10"
	months["Nov"] = "11"
	months["Dec"] = "12"

	timeStcSplit := strings.Split(timeStc, " ")
	year := time.Now().Year()
	yearStr := fmt.Sprint(year)
	month := months[timeStcSplit[0]]

	fmt.Println(yearStr, month)
	date := yearStr + "-" + month
	fmt.Println(date)
}
