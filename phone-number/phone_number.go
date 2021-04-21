package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func filterNumbers(in string) (out string, err error) {
	for _, v := range in {
		switch {
		case unicode.IsDigit(v):
			out += string(v)
		case unicode.IsLetter(v):
			err = errors.New("invalid with letters")
		case v != '-' && v != '(' && v != ')' && v != '.' && unicode.IsPunct(v):
			err = errors.New("invalid with punctuations")
		}
	}
	return
}

func checkLength(in string) (err error) {
	switch {
	case len(in) == 9:
		err = errors.New("invalid when 9 digits")
	case len(in) == 11 && in[0] != '1':
		err = errors.New("invalid when 11 digits does not start with a 1")
	case len(in) > 11:
		err = errors.New("invalid when more than 11 digits")
	}
	return
}

func checkAreaCode(in string) (err error) {
	switch in[0] {
	case '0':
		err = errors.New("invalid if area code starts with 0")
	case '1':
		err = errors.New("invalid if area code starts with 1")
	}
	return
}

func checkExchangeCode(in string) (err error) {
	switch in[3] {
	case '0':
		err = errors.New("invalid if exchange code starts with 0")
	case '1':
		err = errors.New("invalid if exchange code starts with 1")
	}
	return
}

// Number returns formatted phone number
func Number(in string) (out string, err error) {

	if out, err = filterNumbers(in); err != nil {
		return
	}

	if err = checkLength(out); err != nil {
		return
	}

	if len(out) == 11 && out[0] == '1' {
		out = out[1:]
	}

	if err = checkAreaCode(out); err != nil {
		return
	}

	if err = checkExchangeCode(out); err != nil {
		return
	}

	return
}

// AreaCode returns area code for given number
func AreaCode(in string) (string, error) {
	n, err := Number(in)
	var out string
	if err == nil {
		out = n[:3]
	}
	return out, err
}

// Format formats given phone number
func Format(in string) (string, error) {
	n, err := Number(in)
	if err != nil {
		return "", err
	}
	out := fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:])
	return out, err
}
