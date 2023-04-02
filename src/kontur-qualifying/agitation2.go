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

	addressFistDay := parseAddress(firstDaySlsStr)
	fmt.Println()

	buildings := buildNums(nextDaySlsStr, addressFistDay)
	fmt.Println(buildings)
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
	for _, val := range rawAddress {
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

	for key, value := range addresses {
		fmt.Println(key, ":", value)
	}
	return addresses
}

func buildNums(nextDaySlsStr []string, addressFistDay map[string][]int) []int {
	buildings := make([]int, 0)
	for _, street := range nextDaySlsStr {
		_, ok := addressFistDay[street]
		if ok == true {
			num := findBuildNums(addressFistDay[street])
			buildings = append(buildings, num)
			addressFistDay[street] = append(addressFistDay[street], num)
		} else {
			buildings = append(buildings, 1)
			addressFistDay[street] = append(addressFistDay[street], 1)
		}
	}
	return buildings
}

func findBuildNums(buildings []int) int {
	var num int
	for i := 1; i <= len(buildings); i++ {
		var isExists bool
		for _, value := range buildings {
			if value == i {
				isExists = true
				break
			}
		}

		if isExists == false {
			num = i
			fmt.Println("i:", i)
			break
		}
	}
	fmt.Println("returned num:", num)
	return num
}
