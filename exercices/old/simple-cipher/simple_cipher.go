package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

type vigFlag rune

const (
	charStart           = 'a'
	charEnd             = 'z'
	distanceMax         = charEnd - charStart + 1
	vigEncode   vigFlag = 1
	vigDecode   vigFlag = -1
)

var (
	vegKeyRegExp  = regexp.MustCompile(`^[a-z]{3,}$`)
	notWordRegExp = regexp.MustCompile(`[^\w]$`)
)

type shifter struct {
	distance rune
}
type vigener struct {
	key string
}

func addDistance(c, d rune) rune {
	if !unicode.IsLower(c) {
		return 0
	}
	if d < 0 {
		d += distanceMax
	}
	return (c+d-charStart)%distanceMax + charStart
}

func shift(in string, d rune) string {
	out := ""
	for _, c := range strings.ToLower(in) {
		if unicode.IsLower(c) {
			out += string(addDistance(c, d))
		}
	}
	return out
}

func (s shifter) Encode(in string) string {
	return shift(in, s.distance)
}

func (s shifter) Decode(in string) string {
	return shift(in, -s.distance)
}

// NewShift return NewShift cipher
func NewShift(d int) Cipher {
	if d == 0 || d >= distanceMax || d <= -distanceMax {
		return nil
	}
	return shifter{distance: rune(d)}
}

// NewCaesar return NewCaesar cipher
func NewCaesar() Cipher {
	return NewShift(3)
}

func codeVigener(in, key string, flag vigFlag) (out string) {
	count := 0
	for _, c := range strings.ToLower(in) {
		if unicode.IsLower(c) {
			out += string(addDistance(c, rune(flag)*rune(key[count]-'a')))
			count = (count + 1) % len(key)
		}
	}
	return
}

func (v vigener) Encode(in string) string {
	return codeVigener(in, v.key, vigEncode)
}

func (v vigener) Decode(in string) (out string) {
	return codeVigener(in, v.key, vigDecode)
}

// NewVigenere return NewVigenere cipher
func NewVigenere(k string) Cipher {
	if !vegKeyRegExp.MatchString(k) {
		return nil
	}
	return vigener{key: k}
}
