package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	input := "2024-04-13"
	layout := "2006-01-02"

	t, _ := time.Parse(layout, input)

	fmt.Println(t)
	fmt.Println(t.Format("02-Jan-2006"))

	fmt.Println()

	layout2 := "01-02"

	t2, _ := time.Parse(layout2, input)
	fmt.Println(t2)

	sls := []string{"ab cd", "ef"}
	str := "Hey, "
	strCont := str + sls[0]

	fmt.Println(strCont)

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

	inputStr := "Apr 11 23:23:11"
	splittedStr := strings.Split(inputStr, " ")
	monthStr := splittedStr[0]
	monthParsed := months[monthStr]
	fmt.Println(monthParsed)
}
