package strut_meth_interface

import "testing"

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()

	if got != want {
		t.Errorf("got : %.2f want: %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangule := Rectangle{10.0, 10.0}
	got := rectangule.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got : %.2f want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}

func TestCircleArea(t *testing.T) {
	circle := Circle{10}
	want := 314.1592653589793
	checkArea(t, circle, want)
}
