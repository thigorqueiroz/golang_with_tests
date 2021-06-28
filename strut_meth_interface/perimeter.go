package strut_meth_interface

type Rectangule struct {
	width, heigth float64
}

func Perimeter(rectangule Rectangule) float64 {
	return 2 * (rectangule.heigth + rectangule.width)
}

func Area(rectangule Rectangule) float64 {
	return rectangule.width * rectangule.heigth
}
