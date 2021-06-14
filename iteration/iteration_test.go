package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestIteration(t *testing.T) {

	t.Run("a should repeat 5 times", func(t *testing.T) {

		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("got: %q != expected %q ", got, expected)
		}
	})

	t.Run("a double character 5 times, should be 10 character", func(t *testing.T) {
		got := Repeat("aa", 5)
		expected := "aaaaaaaaaa"

		if got != expected {
			t.Errorf("got: %q != expected %q ", got, expected)
		}
	})

}
