package binomial_theorem

func PascalTriangle(n int) []int {
	return binomialExpansion(n, 0)
}

func binomialExpansion(n, k int) []int {
	if k > n {
		return []int{}
	}
	mult := binomialCoefficient(n, k)
	return append([]int{mult}, binomialExpansion(n, k + 1)...)
}

func binomialCoefficient(n, k int) int {
	return Factorial(n) / (Factorial(n - k) * Factorial(k))
}
