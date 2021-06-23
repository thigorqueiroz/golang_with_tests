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
	sums := make([]int, len(numbersToSum))
	for index, arr := range numbersToSum {
		for i := 1; i < len(arr); i++ {
			sums[index] += arr[i]
		}
	}
	return sums
}
