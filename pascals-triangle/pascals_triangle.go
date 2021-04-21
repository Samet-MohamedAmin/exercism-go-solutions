package pascal

var saved = [][]int{
	{1},
}

// Triangle returns pascal triangle
func Triangle(n int) [][]int {
	for i := len(saved); i < n; i++ {
		saved = append(saved, make([]int, i+1))
		saved[i][0], saved[i][i] = 1, 1
		for j := 1; j < i; j++ {
			saved[i][j] = saved[i-1][j-1] + saved[i-1][j]
		}
	}
	return saved[:n]
}
