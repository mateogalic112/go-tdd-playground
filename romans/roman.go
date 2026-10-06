// Package romans with teach us about property based tests
package romans

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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
	{5, "V"},
	{4, "IV"},
	{1, "I"},
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

func ConvertToArabicRecursive(roman string, arabic int) (int, error) {
	if len(roman) == 0 {
		return arabic, nil
	}

	for _, numeral := range allRomanNumerals {
		if after, found := strings.CutPrefix(roman, numeral.Symbol); found {
			return ConvertToArabicRecursive(after, arabic+numeral.Value)
		}
	}

	return 0, errors.New("invalid string")

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
