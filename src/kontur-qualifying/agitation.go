package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	firstDaySlsStr, nextDaySlsStr := fillData()
	fmt.Println(firstDaySlsStr)
	fmt.Println(nextDaySlsStr)

	addressFistDay := parseAddress(firstDaySlsStr)

}

func fillData() (firstDaySlsStr, nextDaySlsStr []string) {
	firstDaySlsStr = make([]string, 0)
	nextDaySlsStr = make([]string, 0)

	var numPrevDay, numNextDay int
	reader := bufio.NewReader(os.Stdin)
	var elem string

	fmt.Fscan(reader, &numPrevDay)
	for i := 0; i < numPrevDay; i++ {
		fmt.Fscan(reader, &elem)
		firstDaySlsStr = append(firstDaySlsStr, elem)
	}

	fmt.Fscan(reader, &numNextDay)
	for i := 0; i < numNextDay; i++ {
		fmt.Fscan(reader, &elem)
		nextDaySlsStr = append(nextDaySlsStr, elem)
	}

	return firstDaySlsStr, nextDaySlsStr
}

func parseAddress(rawAddress []string) map[string][]int {
	addresses := make(map[string][]int)
	for i, val := range rawAddress {
		var streetRune []rune
		var buildingRune []rune
		runeAddress := []rune(val)
		for i := 0; i < len(runeAddress); i++ {
			digit := unicode.IsDigit(runeAddress[i])
			if digit == true {
				buildingRune = append(buildingRune, runeAddress[i])
			} else {
				streetRune = append(streetRune, runeAddress[i])
			}
		}

		streetStr := string(streetRune[:])
		buildingStr := string(buildingRune[:])
		buildingInt, _ := strconv.Atoi(buildingStr)
		addresses[streetStr] = append(addresses[streetStr], buildingInt)
	}

	return addresses
}
