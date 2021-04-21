package allergies

var list = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func isAllergic(score uint, i int) bool {
	return score>>i%2 == 1
}

// Allergies returns allergies
func Allergies(score uint) (out []string) {
	for i := range list {
		if isAllergic(score, i) {
			out = append(out, list[i])
		}
	}
	return
}

func findSubstanceIndex(substance string) (i int, ok bool) {
	for i = range list {
		if list[i] == substance {
			ok = true
			return
		}
	}
	return
}

// AllergicTo returns true if allergic to
func AllergicTo(score uint, substance string) bool {
	i, found := findSubstanceIndex(substance)
	if found && isAllergic(score, i) {
		return true
	}
	return false
}
