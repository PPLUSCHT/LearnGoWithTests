package ch5structs

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return 2.0 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
