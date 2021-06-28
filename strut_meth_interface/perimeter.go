package strut_meth_interface

import "math"

type Rectangule struct {
	width, heigth float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (r Rectangule) area() float64 {
	return r.heigth * r.width
}

func (r Rectangule) perimeter() float64 {
	return 2 * (r.heigth + r.width)
}
