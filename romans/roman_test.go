package romans

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

var mixedCases = []struct {
	Roman  string
	Arabic uint16
	Valid  bool
}{
	// Basic numerals

	{Roman: "I", Arabic: 1, Valid: true},
	{Roman: "V", Arabic: 5, Valid: true},
	{Roman: "X", Arabic: 10, Valid: true},
	{Roman: "L", Arabic: 50, Valid: true},
	{Roman: "C", Arabic: 100, Valid: true},
	{Roman: "D", Arabic: 500, Valid: true},
	{Roman: "M", Arabic: 1000, Valid: true},
	// Repetition
	{Roman: "II", Arabic: 2, Valid: true},
	{Roman: "III", Arabic: 3, Valid: true},
	{Roman: "XX", Arabic: 20, Valid: true},
	{Roman: "XXX", Arabic: 30, Valid: true},
	{Roman: "CCC", Arabic: 300, Valid: true},
	{Roman: "MMM", Arabic: 3000, Valid: true},
	// Additive notation
	{Roman: "VI", Arabic: 6, Valid: true},
	{Roman: "VII", Arabic: 7, Valid: true},
	{Roman: "VIII", Arabic: 8, Valid: true},
	{Roman: "XII", Arabic: 12, Valid: true},
	{Roman: "LX", Arabic: 60, Valid: true},
	{Roman: "CL", Arabic: 150, Valid: true},
	{Roman: "MD", Arabic: 1500, Valid: true},
	// Subtractive notation
	{Roman: "IV", Arabic: 4, Valid: true},
	{Roman: "IX", Arabic: 9, Valid: true},
	{Roman: "XL", Arabic: 40, Valid: true},
	{Roman: "XC", Arabic: 90, Valid: true},
	{Roman: "CD", Arabic: 400, Valid: true},
	{Roman: "CM", Arabic: 900, Valid: true},
	// Mixed / realistic values
	{Roman: "XIV", Arabic: 14, Valid: true},
	{Roman: "XIX", Arabic: 19, Valid: true},
	{Roman: "XLII", Arabic: 42, Valid: true},
	{Roman: "XCIX", Arabic: 99, Valid: true},
	{Roman: "CDXLIV", Arabic: 444, Valid: true},
	{Roman: "CMXCIX", Arabic: 999, Valid: true},
	{Roman: "MCMXCIV", Arabic: 1994, Valid: true},
	{Roman: "MM", Arabic: 2000, Valid: true},
	{Roman: "MMXXVI", Arabic: 2026, Valid: true},
	{Roman: "MMMDCCCLXXXVIII", Arabic: 3888, Valid: true},
	{Roman: "MMMCMXCIX", Arabic: 3999, Valid: true},
	// Invalid characters
	{Roman: "Q", Arabic: 0, Valid: false},
	{Roman: "QQ", Arabic: 0, Valid: false},
	{Roman: "ABC", Arabic: 0, Valid: false},
	{Roman: "XQ", Arabic: 0, Valid: false},
	{Roman: "123", Arabic: 0, Valid: false},
	// Too many repetitions
	{Roman: "IIII", Arabic: 0, Valid: false},
	{Roman: "XXXX", Arabic: 0, Valid: false},
	{Roman: "CCCC", Arabic: 0, Valid: false},
	{Roman: "MMMM", Arabic: 0, Valid: false},
	// V, L and D cannot repeat
	{Roman: "VV", Arabic: 0, Valid: false},
	{Roman: "LL", Arabic: 0, Valid: false},
	{Roman: "DD", Arabic: 0, Valid: false},
	// Invalid subtraction
	{Roman: "IL", Arabic: 0, Valid: false},
	{Roman: "IC", Arabic: 0, Valid: false},
	{Roman: "ID", Arabic: 0, Valid: false},
	{Roman: "IM", Arabic: 0, Valid: false},
	{Roman: "XD", Arabic: 0, Valid: false},
	{Roman: "XM", Arabic: 0, Valid: false},
	{Roman: "VX", Arabic: 0, Valid: false},
	{Roman: "LC", Arabic: 0, Valid: false},
	{Roman: "DM", Arabic: 0, Valid: false},
	// Invalid ordering / non-canonical forms
	{Roman: "IIV", Arabic: 0, Valid: false},
	{Roman: "IIX", Arabic: 0, Valid: false},
	{Roman: "XXC", Arabic: 0, Valid: false},
	{Roman: "CCM", Arabic: 0, Valid: false},
	{Roman: "IXIX", Arabic: 0, Valid: false},
	{Roman: "IVIV", Arabic: 0, Valid: false},
	{Roman: "CMCM", Arabic: 0, Valid: false},
	{Roman: "VIV", Arabic: 0, Valid: false},
	{Roman: "XCM", Arabic: 0, Valid: false},
	// Ordering violations
	{Roman: "IXX", Arabic: 0, Valid: false},
	{Roman: "XIXI", Arabic: 0, Valid: false},
	{Roman: "IVI", Arabic: 0, Valid: false},
	{Roman: "CMC", Arabic: 0, Valid: false},
	// Edge cases
	{Roman: "", Arabic: 0, Valid: false},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got, _ := ConvertToArabicRecursive(test.Roman, 0, nil)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
		Values: func(args []reflect.Value, r *rand.Rand) {
			args[0] = reflect.ValueOf(uint16(r.Intn(4000)))
		},
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestInvalidConvertingToArabic(t *testing.T) {
	for _, test := range mixedCases {
		t.Run(fmt.Sprintf("%q should be %t", test.Roman, test.Valid), func(t *testing.T) {
			got, err := ConvertToArabicRecursive(test.Roman, 0, nil)
			if err != nil {
				if test.Valid {
					t.Errorf("roman string %q should have been invalid", test.Roman)
				}
			}

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}
