package romannumerals

import (
	"errors"
)

var m = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func digitToRomanNumeral(digit, ten int) string {
	out := ""
	if digit < 1 || digit > 9 {
		return out
	}
	if digit >= 4 {
		key := (digit + 1) / 5
		out = m[ten*key*5]
	}
	if digit%5 == 4 {
		return m[ten] + out
	}
	for i := 0; i < digit%5; i++ {
		out += m[ten]
	}
	return out
}

// ToRomanNumeral converts arabic number to roman number
func ToRomanNumeral(n int) (out string, err error) {
	if n < 1 || n > 3000 {
		err = errors.New("roman numerals does not support numbers strictly less than 1 or superior than 3000")
		return
	}
	for i := 1; i <= n; i *= 10 {
		digit := (n / i) % 10
		out = digitToRomanNumeral(digit, i) + out
	}
	return
}
