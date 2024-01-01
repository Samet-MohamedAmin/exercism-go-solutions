package binarysearch

// SearchInts performs binary search on given sorted slice for given key
// and returns index of key if exists in the given slice or -1 if not
func SearchInts(slice []int, key int) int {
	left := 0
	right := len(slice) - 1

	for {
		if left > right {
			return -1
		}

		mid := (left + right) / 2

		switch key {
		case slice[left]:
			return left
		case slice[mid]:
			return mid
		case slice[right]:
			return right
		}

		if key < slice[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}
