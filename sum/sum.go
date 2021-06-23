package sum

func Sum(numbers []int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, num := range numbersToSum {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) <= 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
