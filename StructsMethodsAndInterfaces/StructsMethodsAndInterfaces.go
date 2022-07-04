package StructsMethodsAndInterfaces

import "math"

// To implement an interface in Go, we just need to implement all the methods in the interface.
type Shape interface {
	// Area Wherever Area method is present, interface will be implemented.
	Area() float64
}

// A Struct is just a named collection of fields where data can be stored.
type Rectangle struct {
	Height, Width float64
}

// Define a method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circles struct {
	Radius float64
}

func (c Circles) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Height, Width float64
}

// If we comment out this method which is related to Triangle Struct then this
// will not be considered as an interface since Area() is not implemented.
// And in interface we have defined the Area() method.

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Width
}

func Perimeter(Height float64, Width float64) float64 {
	return 2 * (Height + Width)
	// return 0
}

func Area(Height float64, Width float64) float64 {
	return Height * Width
}

// Refactoring for struct
func PerimeterV2(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
	// return 0
}

func AreaV2(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
