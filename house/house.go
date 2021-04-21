package house

import (
	"fmt"
	"strings"
)

const (
	firstStart = "This"
	firstVerb  = "is"
	otherStart = "that"
)

var verbs = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

var objects = []string{
	"house that Jack built",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var savedParts = []string{}

func getLine(start, verb, object string) string {
	return fmt.Sprintf("%s %s the %s", start, verb, object)
}

func fillSavedParts(n int) {
	for i := len(savedParts); i < n; i++ {
		savedParts = append(savedParts, getLine(otherStart, verbs[i], objects[i]))
	}
}

func reverseSavedParts(n int) (out []string) {
	for i := n - 1; i >= 0; i-- {
		out = append(out, savedParts[i])
	}
	return
}

// Verse returns verse
func Verse(n int) string {
	fillSavedParts(n - 1)
	firstLine := getLine(firstStart, firstVerb, objects[n-1])
	l := append([]string{firstLine}, reverseSavedParts(n-1)...)
	return strings.Join(l, "\n") + "."
}

// Song returns song
func Song() string {
	l := []string{}
	for i := 1; i < len(objects)+1; i++ {
		l = append(l, Verse(i))
	}
	return strings.Join(l, "\n\n")
}
