package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expected := 4
	if got != expected {
		t.Errorf("'GOT' is diferent from 'EXPECTED'. got : %q | expected : %q")
	}
}
