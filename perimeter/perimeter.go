package main

import "math"

// method of rectangle and circle are satisfied with interface output
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// method of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// method of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
