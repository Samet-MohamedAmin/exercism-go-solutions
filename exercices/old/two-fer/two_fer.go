package twofer

import "fmt"

//ShareWith returns the result of the ShareWith task
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}

	result := fmt.Sprintf("One for %s, one for me.", name)

	return result

}
