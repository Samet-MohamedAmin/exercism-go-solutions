//Package hamming calculates the Hamming difference between two DNA strands
package hamming

import "errors"

//Distance calculates the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands do not have same length")
	}

	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
