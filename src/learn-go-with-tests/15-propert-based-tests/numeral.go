// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals
package roman

import (
	"strings"
)

type RomanNumerals struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {

	var arabic = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
