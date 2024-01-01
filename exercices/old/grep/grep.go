package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// flags:
	// - `-n` Print the line numbers of each matching line.
	// - `-l` Print only the names of files that contain at least one matching line.
	// - `-i` Match line using a case-insensitive comparison.
	// - `-v` Invert the program -- collect all lines that fail to match the pattern.
	// - `-x` Only match entire lines, instead of lines that contain a match.
	flagPrintLineNumbers      = "-n"
	flagPrintOnlyNamesOfFiles = "-l"
	flagCaseInsensitive       = "-i"
	flagInvertMatch           = "-v"
	flagEntireLine            = "-x"
)

func searchFile(pattern string, file string, mapFlags map[string]bool) ([]string, bool) {
	out := []string{}
	p := pattern

	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// convert pattern to lower case
	if mapFlags[flagCaseInsensitive] {
		p = strings.ToLower(pattern)
	}

	// iterate over file lines
	i := 0
	for scanner.Scan() {
		i++
		lineOriginal := scanner.Text()
		line := lineOriginal

		// convert line to lower case
		if mapFlags[flagCaseInsensitive] {
			line = strings.ToLower(line)
		}
		// check if line contains search pattern
		found := false
		if mapFlags[flagEntireLine] {
			found = p == line
		} else {
			found = strings.Contains(line, p)
		}
		// add search result
		if (!mapFlags[flagInvertMatch]) && found || mapFlags[flagInvertMatch] && (!found) {
			var l string
			switch {
			case mapFlags[flagPrintOnlyNamesOfFiles]:
				return out, true
			case mapFlags[flagPrintLineNumbers]:
				l = fmt.Sprintf("%d:%s", i, lineOriginal)
			default:
				l = lineOriginal
			}
			out = append(out, l)
		}
	}
	return out, len(out) > 0
}

// Search returns search result for given pattern, flags and files
func Search(pattern string, flags []string, files []string) []string {
	out := []string{}

	// map flags
	mapFlags := make(map[string]bool, len(flags))
	for _, f := range flags {
		mapFlags[f] = true
	}

	for _, file := range files {
		r, found := searchFile(pattern, file, mapFlags)
		switch {
		case mapFlags[flagPrintOnlyNamesOfFiles] && found:
			out = append(out, file)
		case len(files) == 1:
			out = append(out, r...)
		case len(files) > 1:
			for i := range r {
				e := fmt.Sprintf("%s:%s", file, r[i])
				out = append(out, e)
			}
		}
	}

	return out
}
