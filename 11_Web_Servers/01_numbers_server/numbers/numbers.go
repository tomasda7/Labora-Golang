package numbers

func SumFirstOdds(n int) int {
	var acc int
	odd := 1

	for i := 0; i < n; i++ {
		acc += odd
		odd += 2
	}

	return acc
}

func SumFirstEvens(n int) int {
	return n * (n + 1)
}
