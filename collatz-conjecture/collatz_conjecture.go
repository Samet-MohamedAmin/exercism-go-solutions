package collatzconjecture

import "errors"

// CollatzConjecture returns CollatzConjecture result
func CollatzConjecture(in int) (steps int, err error) {
	switch {
	case in == 1:
		break
	case in == 0:
		err = errors.New("zero is an error")
	case in < 0:
		err = errors.New("negative value is an error")
	case in%2 == 0:
		steps, err = CollatzConjecture(in / 2)
		steps++
	case in%2 == 1:
		steps, err = CollatzConjecture(3*in + 1)
		steps++
	}
	return
}
