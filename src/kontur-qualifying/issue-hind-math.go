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
	fmt.Println(dif)
}

func digitsPremutation(num int) int {
	// Конвертируем входные данные в строку для парсинга по символам
	numStr := strconv.Itoa(num)

	// Парсим входные данные, в целочисленный (интовый) слайс numInt []int
	var numInt []int
	for _, value := range numStr {
		strVal := string(value)
		intVal, _ := strconv.Atoi(strVal)
		numInt = append(numInt, intVal)
	}

	// Готовим int слайсы (каждый элемент слайса - отдельная цифра входных данных)
	// для дальнейшей сортировки
	maxIntSlice := make([]int, len(numInt))
	minIntSlice := make([]int, len(numInt))
	copy(maxIntSlice[0:], numInt[0:])
	copy(minIntSlice, numInt[0:])

	// Сортируем слайс по убыванию для комбинирования максимально возможного значения
	sort.Slice(maxIntSlice, func(i, j int) bool {
		return maxIntSlice[i] > maxIntSlice[j]
	})

	// Сортируем слайс по возрастанию для комбинирования минимально возможного значения
	sort.Slice(minIntSlice, func(i, j int) bool {
		return minIntSlice[i] < minIntSlice[j]
	})

	// Переводим результат сортировки в строку (string) для дальнейшего
	// конвертирования в целочиленное значение (int)
	var maxIntStr string
	var minIntStr string

	for _, value := range maxIntSlice {
		maxIntStr += strconv.Itoa(value)
	}

	for _, value := range minIntSlice {
		minIntStr += strconv.Itoa((value))
	}

	// Конвертируем string в int
	maxInt, err := strconv.Atoi(maxIntStr)
	if err != nil {
		fmt.Println("converting from string to int: ", err.Error())
	}

	minInt, err := strconv.Atoi(minIntStr)
	if err != nil {
		fmt.Println("converting from string to int: ", err.Error())
	}

	// Вычисляем разницу между максимальным и минимальным значением
	dif := maxInt - minInt

	// Возвращаем результат
	return dif
}
