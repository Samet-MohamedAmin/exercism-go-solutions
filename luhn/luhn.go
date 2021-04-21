package luhn

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Transform removes white spaces and returns an error if the input
// contains charcters other than digits or white spaces
func Transform(input string) (string, error) {

	result := ""
	for _, char := range input {
		if char == ' ' {
			continue
		}
		if unicode.IsNumber(char) {
			result += string(char)
			continue
		}

		return "", errors.New("charcters other than digits and white spaces are not allowed")
	}

	return result, nil
}

// Sum returns the luhn sum for a proper credit card input.
// A proper credit card input should only contain digits and white spaces.
func Sum(input string) (int, error) {

	sum := 0
	for i := 1; i <= len(input); i++ {
		// Iterate input characters one by one from right to left
		c := input[len(input)-i]
		// covert character to int
		value, err := strconv.Atoi(string(c))
		if err != nil {
			return -1, err
		}
		if i%2 == 0 {
			if value != 9 {
				value *= 2
				value %= 9
			}
		}
		sum += value
	}

	return sum, nil
}

// Valid checks if the input is a valid credit card number
// using the luhn method
func Valid(input string) bool {

	input, err := Transform(input)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if len(input) < 2 {
		fmt.Println("input length should be strictly greater than 2")
		return false
	}

	sum, _ := Sum(input)
	if sum%10 != 0 {
		return false
	}

	return true
}
