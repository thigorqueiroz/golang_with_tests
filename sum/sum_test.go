package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("GOT %d is diferent from WANT %d, value: %v", got, want, numbers)
		}
	})

	t.Run("sum of all numbers", func(t *testing.T) {
		numbersA := []int{1, 2, 3}
		numbersB := []int{2, 3, 4}
		got := SumAll(numbersA, numbersB)
		want := []int{6, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("GOT %d  WANT %d", got, want)
		}
	})

	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{2, 4, 5, 6})
		want := []int{9, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}

	})

}

func TestSlice(t *testing.T) {
	sliceSample := []int{1, 2, 3, 4, 5, 6}
	got := sliceSample[3:]
	want := []int{4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d , want %d ", got, want)
	}

}
