package diffsquares


func SquareOfSum(n int) int {
	var sum int 
	for num := 1; num <= n; num++ {
		sum += num
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for num := 1; num <= n; num++ {
		sum += num * num
	}
	return sum
}

func Difference(n int) int {
	var squareOfSum = SquareOfSum(n)
	var sumOfSquares = SumOfSquares(n)
	return squareOfSum - sumOfSquares
}
