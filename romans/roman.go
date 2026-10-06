// Package romans with teach us about property based tests
package romans

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value      int
	Symbol     string
	MaxRepeats int
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M", 3},
	{900, "CM", 1},
	{500, "D", 1},
	{400, "CD", 1},
	{100, "C", 3},
	{90, "XC", 1},
	{50, "L", 1},
	{40, "XL", 1},
	{10, "X", 3},
	{9, "IX", 1},
	{5, "V", 1},
	{4, "IV", 1},
	{1, "I", 3},
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

func ConvertToArabicRecursive(roman string, arabic int, processed []RomanNumeral) (int, error) {
	if len(roman) == 0 {
		return arabic, nil
	}

	for _, numeral := range allRomanNumerals {
		if after, found := strings.CutPrefix(roman, numeral.Symbol); found {
			if len(processed) == 0 {
				return ConvertToArabicRecursive(after, arabic+numeral.Value, append(processed, numeral))
			}

			lastProcessed := processed[len(processed)-1]
			if lastProcessed.Value < numeral.Value {
				return 0, errors.New("preceeding symbol cant be smaller than next one")
			}

			if len(processed) < numeral.MaxRepeats {
				return ConvertToArabicRecursive(after, arabic+numeral.Value, append(processed, numeral))
			}

			var counter = 0
			for _, item := range processed[:numeral.MaxRepeats] {
				if item.Symbol == numeral.Symbol {
					counter++
				}
			}

			if counter == numeral.MaxRepeats {
				return 0, errors.New("too many items in sequence")
			}

			return ConvertToArabicRecursive(after, arabic+numeral.Value, append(processed, numeral))
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
