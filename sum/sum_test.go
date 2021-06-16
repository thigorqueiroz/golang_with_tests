package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("sum of numbers using slice as parameter", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("GOT %d is diferent from WANT %d, value: %v", got, want, numbers)
		}
	})

}
