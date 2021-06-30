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
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, heigth: 6}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, want: 31.0},
	}

	for _, tt := range areaTest {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got: %g | want %g", tt.shape, got, tt.want)
			}
		})
	}

}

func TestCircleArea(t *testing.T) {
	circle := Circle{10}
	want := 314.1592653589793
	checkArea(t, circle, want)
}
