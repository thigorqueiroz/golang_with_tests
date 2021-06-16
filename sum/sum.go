package sum

func Sum(numbers []int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}
