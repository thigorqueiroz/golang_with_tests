package strut_meth_interface

import "math"

type Rectangle struct {
	width, heigth float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.heigth * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.heigth + r.width)
}
