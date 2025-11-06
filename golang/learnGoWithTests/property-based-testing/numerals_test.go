package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertingToRoman(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("Got: %q, Want: %q", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%s gets converted to %d", test.Roman, test.Arabic)
		t.Run(description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("Got: %d, Want: %d", got, want)
			}
		})
	}
}

func TestConvertingToArabicRecursive(t *testing.T) {
	for _, test := range cases {
		description := fmt.Sprintf("%s gets converted to %d", test.Roman, test.Arabic)
		t.Run(description, func(t *testing.T) {
			got := ConvertToArabicRecursive(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("Got: %d, Want: %d", got, want)
			}
		})
	}
}

func TestNum(t *testing.T) {
	fmt.Println(ConvertToRoman(uint16(9999)))
}

var bigNum = "MMMMMMMMMMMMMMMMMMMMMMMMMMMCDXXIV"

func BenchmarkConvertArabic(b *testing.B) {
	for b.Loop() {
		ConvertToArabic(bigNum)
	}
}

func BenchmarkConvertArabicRecursive(b *testing.B) {
	for b.Loop() {
		ConvertToArabicRecursive(bigNum)
	}
}

func TestPropertiesOfRecursiveImpl(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)

		return ConvertToArabic(roman) == ConvertToArabicRecursive(roman)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
