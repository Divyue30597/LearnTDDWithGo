package StructsMethodsAndInterfaces

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

