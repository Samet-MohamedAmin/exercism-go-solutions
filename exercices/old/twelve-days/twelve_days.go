package twelve

import "fmt"

var verseStructure = []struct {
	nth     string
	special string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

const (
	baseString = "On the %s day of Christmas my true love gave to me: %s."
)

// Verse returns single verse
func Verse(input int) string {
	lastPart := verseStructure[0].special
	for i := 1; i < input; i++ {
		if i == 1 {
			lastPart = "and " + lastPart
		}
		lastPart = fmt.Sprintf("%s, %s", verseStructure[i].special, lastPart)
	}

	return fmt.Sprintf(baseString, verseStructure[input-1].nth, lastPart)
}

// Song returns the entire song
func Song() string {
	result := Verse(1)
	for i := 2; i <= len(verseStructure); i++ {
		result += "\n" + Verse(i)
	}
	return result
}
