package ch4slices

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, numbers := range slices {
		if len(numbers) >= 2 {
			sums[i] = Sum(numbers[1:])
		} else {
			sums[i] = 0
		}
	}
	return sums
}
