package sum

import "fmt"

func Sum(numbers []int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbers ...[]int) (sum []int) {
	for index, arr := range numbers {
		var sumOfThis int = 0
		for _, number := range arr {
			sumOfThis += number
		}
		fmt.Printf("index: %d", index)
		//sumOfThis = 0
	}
	return
}
