package strand

var complements = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns dna complement
func ToRNA(dna string) (res string) {

	for _, c := range dna {
		res += string(complements[c])
	}
	return
}
