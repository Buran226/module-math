package math

func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Exponential(n int, i int) int {
	res := n
	for dt := 1; dt < i; dt++ {
		res = int(n * res)
	}
	return res
}

func SuperMult(args ...int) int {
	res := 1
	for _, arg := range args {
		res *= arg
	}
	return res
}
