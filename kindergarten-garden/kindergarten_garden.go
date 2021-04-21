package kindergarten

import (
	"errors"
	"regexp"
	"sort"
)

// Garden holds garden variable
type Garden map[string][]string

var (
	mapPlant = map[byte]string{
		'C': "clover",
		'G': "grass",
		'R': "radishes",
		'V': "violets",
	}
	r = regexp.MustCompile(`^\n(\w+)\n(\w+)$`)

	errBadDiagram    = errors.New("bad diagram")
	errChildrenEmpty = errors.New("there are no children")
	errDuplicateName = errors.New("duplicate name")
)

func checkDiagramLines(lines []string) bool {
	if len(lines) == 0 {
		return false
	}
	for _, l := range lines[1:] {
		if len(lines[0]) != len(l) || len(l)%2 != 0 {
			return false
		}
	}
	return true
}

func getLines(diagram string) (out []string) {
	if out = r.FindStringSubmatch(diagram); len(out) > 0 {
		out = out[1:]
	}
	return
}

func genChildPlants(lines []string, index, step int) (out []string) {
	start := index * step
	end := (index + 1) * step
	for _, l := range lines {
		for i := start; i < end; i++ {
			out = append(out, mapPlant[l[i]])
		}
	}
	return out
}

// NewGarden returns new Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(children) == 0 {
		return nil, errChildrenEmpty
	}

	childrenSorted := append([]string{}, children...)
	sort.Strings(childrenSorted)
	if len(childrenSorted) != len(children) {
		return nil, errDuplicateName
	}

	lines := getLines(diagram)
	if len(lines) < 2 {
		return nil, errBadDiagram
	}
	if !checkDiagramLines(lines) {
		return nil, errBadDiagram
	}

	g := Garden{}
	step := len(lines[0]) / len(childrenSorted)
	for i, c := range childrenSorted {
		g[c] = genChildPlants(lines, i, step)
	}
	return &g, nil
}

// Plants returns garden plants
func (g *Garden) Plants(child string) ([]string, bool) {
	out, ok := (*g)[child]
	return out, ok
}
