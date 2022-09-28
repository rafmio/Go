package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func convFahToCel(fahrenheit *float64) float64 {
	return (*fahrenheit - 32) * 5 / 9
}

func main() {
	fmt.Print("Enter a tem in Fahrenheit: ")
	fahrenheit, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := convFahToCel(&fahrenheit)
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
