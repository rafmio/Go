package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	firstDaySlsStr, nextDaySlsStr := fillData()

	addressFistDay := parseAddress(firstDaySlsStr)
	fmt.Println()

	buildings := findBuildNums(nextDaySlsStr, addressFistDay)
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

func findBuildNums(nextDaySlsStr []string, addressFistDay map[string][]int) []int {
	buildings := make([]int, 0)            // создаем слайс для вывода номеров домов
	for _, street := range nextDaySlsStr { // перебираем улицы второго дня
		_, ok := addressFistDay[street] // проверяем улицу на предмет наличия в первом дне
		if ok == true {                 // если улица уже была в предыдущем дне...
			sort.Ints(addressFistDay[street])
			fmt.Println("sorted:", street, ":", addressFistDay[street])

			for _, buildNum := range addressFistDay[street] {
				for i := 1; i < len(addressFistDay[street]); i++ {
					if buildNum == i {
						break
					} else {
						buildings = append(buildings, i)
						addressFistDay[street] = append(addressFistDay[street], i)
						break
					}
				}
				break
			}
			//------------------------
			// for i := 1; i <= len(addressFistDay[street]); i++ { // перебираем номера домов от 1
			// 	for _, buildNum := range addressFistDay[street] { // перебираем уже расклеенные номера домов
			// 		if buildNum == i { // если при переборе текущего дома уже есть такое номер
			// 			continue // пропускаем
			// 		} else { // если такого дома еще нет
			// 			addressFistDay[street] = append(addressFistDay[street], i) // добавляем в слайс уже расклеенных домов и
			// 			fmt.Println("appended i:", i)
			// 			buildings = append(buildings, i) // и добавляем в слайс на вывод
			// 			break                            // прерываем перебор слайса с расклееными домами
			// 		}
			// 	}

			// 	break // прерываем перебор номеров
			// }
			//------------------------
		} else { // если такой улицы еще не было, то
			addressFistDay[street] = append(addressFistDay[street], 1) // добавляем 1й дом в раклееные номера
			buildings = append(buildings, 1)                           // и добавляем 1й дом в слайс на вывод

		}
	}

	fmt.Println("inside finidBuildNums: ", buildings)
	return buildings
}
