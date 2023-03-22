/*
Входные данные - целочисленное значение (int), не содержащее нулей

Задача:
Скомбинировать из цифр значения максимально и минимально возможные значения.
Вычислить разницу

Выходные данные - целочисленное значени (int) междуу максимально и минимально
возможными значениями из комбинаций цифр
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Эмулируем вход данных и вызоваем функцию
	var num int = 4965
	dif := digitsPremutation(num)

	// Выводим результат
	fmt.Println("dif: ", dif)
}

func digitsPremutation(num int) int {
	numStr := strconv.Itoa(num)

	var numInt []int
	for _, value := range numStr {
		strVal := string(value)
		intVal, _ := strconv.Atoi(strVal)
		numInt = append(numInt, intVal)
	}

	for i, val := range numInt {
		fmt.Println(i, " : ", val)
	}
	fmt.Println()

	maxIntSlice := make([]int, len(numInt))
	minIntSlice := make([]int, len(numInt))
	copy(maxIntSlice[0:], numInt[0:])
	copy(minIntSlice, numInt[0:])

	sort.Slice(maxIntSlice, func(i, j int) bool {
		return maxIntSlice[i] > maxIntSlice[j]
	})
	fmt.Println("maxIntSlie: ", maxIntSlice)

	sort.Slice(minIntSlice, func(i, j int) bool {
		return minIntSlice[i] < minIntSlice[j]
	})

	fmt.Println(maxIntSlice)
	fmt.Println(minIntSlice)

	var maxIntStr string
	var minIntStr string

	for _, value := range maxIntSlice {
		maxIntStr += strconv.Itoa(value)
	}

	for _, value := range minIntSlice {
		minIntStr += strconv.Itoa((value))
	}
	fmt.Println(maxIntStr)
	fmt.Println(minIntStr)

	maxInt, err := strconv.Atoi(maxIntStr)
	if err != nil {
		fmt.Println("converting from string to int: ", err.Error())
	}

	minInt, err := strconv.Atoi(minIntStr)
	if err != nil {
		fmt.Println("converting from string to int: ", err.Error())
	}
	fmt.Printf("maxInt: %d, minInt: %d, type of maxInt: %T, type of minInt: %T\n",
		maxInt, minInt, maxInt, minInt,
	)

	dif := maxInt - minInt

	return dif
}
