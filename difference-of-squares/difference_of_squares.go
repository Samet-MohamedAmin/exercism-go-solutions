package diffsquares

//SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

//SumOfSquares returns the sum the squares of the first n natural numbers
func SumOfSquares(n int) int {

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

//Difference returns diffrence between SquareOfSum and SumOfSquares
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)
}
