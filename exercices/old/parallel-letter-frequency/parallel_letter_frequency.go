package letter

import "fmt"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Merge input FreqMap into 1 FreqMap that contains the sum of the FreqMap input
func Merge(am ...FreqMap) FreqMap {
	merged := FreqMap{}
	for _, m := range am {
		for k, v := range m {
			merged[k] += v
		}
	}

	return merged
}

// ConcurrentFrequency counts the frequency of each rune in a given array of strings and returns this
// data as a FreqMap. This is done in a concurrent way
func ConcurrentFrequency(as []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range as {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	k := cap(c)
	fmt.Println(k)

	result := FreqMap{}
	for range as {
		result = Merge(result, <-c)
	}

	return result
}
