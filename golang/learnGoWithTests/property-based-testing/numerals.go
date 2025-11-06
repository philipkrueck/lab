package numerals

import (
	"strings"
)

type RomanNumeral struct {
	Arabic uint16
	Roman  string
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, romanNumeral := range allRomanNumerals {
		for arabic >= romanNumeral.Arabic {
			result.WriteString(romanNumeral.Roman)
			arabic -= romanNumeral.Arabic
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (result uint16) {
	for _, romanNumeral := range allRomanNumerals {
		for strings.HasPrefix(roman, romanNumeral.Roman) {
			result += romanNumeral.Arabic
			roman = strings.TrimPrefix(roman, romanNumeral.Roman)
		}
	}

	return
}

func ConvertToArabicRecursive(roman string) uint16 {
	for _, romanNumeral := range allRomanNumerals {
		if remaining, hasPrefix := strings.CutPrefix(roman, romanNumeral.Roman); hasPrefix {
			return romanNumeral.Arabic + ConvertToArabic(remaining)
		}
	}

	return 0
}
