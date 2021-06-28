package strut_meth_interface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangule := Rectangule{10.0, 10.0}
	got := rectangule.perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got : %.2f want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangule := Rectangule{12.0, 6.0}
	got := rectangule.area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{10}
	got := circle.area()
	want := 314.1592653589793

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
