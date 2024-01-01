package raindrops

import "fmt"

var dropMap = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

//Convert generates raindrops sound from drop number input
func Convert(drop int) string {
	result := ""

	for number, sound := range dropMap {
		if drop%number == 0 {
			result += sound
		}
	}

	if len(result) == 0 {
		return fmt.Sprint(drop)
	}

	return result
}
