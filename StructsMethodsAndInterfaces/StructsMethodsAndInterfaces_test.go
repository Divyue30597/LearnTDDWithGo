package StructsMethodsAndInterfaces

import (
	"testing"
)

// Test for the perimeter
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		// f is for float and .2 means print upto 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// V2 using structs
func TestPerimeterV2(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := PerimeterV2(rectangle)
	want := 40.0

	if got != want {
		// f is for float and .2 means print upto 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaV2(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := AreaV2(rectangle)
	want := 72.0

	if got != want {
		// Use of g will print a more precise decimal number in the error message
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaV3(t *testing.T) {

	// Area for rectangles
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := AreaV2(rectangle)
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	// Area for circles
	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circles{10}
	// We will get this issue while using AreaV2 with the circle -
	// cannot use circle (variable of type Circles) as type Rectangle in argument to AreaV2
	// What can be done?? -> 2 options
	// You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.
	// We can define methods on our newly defined types instead.
	// got := AreaV2(circle)
	// want := 314.1592653589793

	// if got != want {
	// 	t.Errorf("Got %g Want %g", got, want)
	// }
}

// Introducing Methods:
// We need methods since golang does not have classes. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// A method is a function with a receiver.
// Declaring a method : func (receiverName receiverType) MethodName(Args) {}
func TestAreaV4(t *testing.T) {

	// Area for rectangles
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	// Area for circles
	t.Run("circles", func(t *testing.T) {
		circle := Circles{10}
		// We will get this issue while using AreaV2 with the circle -
		// cannot use circle (variable of type Circles) as type Rectangle in argument to AreaV2
		// What can be done?? -> 2 options
		// You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.
		// We can define methods on our newly defined types instead.
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Got %g Want %g", got, want)
		}
	})
}

// Refactoring further more for the interfaces
// Interfaces - They are very powerful concept in statically typed languages because allow to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.
func TestAreaV5(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circles := Circles{10}
		checkArea(t, circles, 314.1592653589793)
	})
}

// We can further refactor our tests to Table driven tests

func TestAreaV6(t *testing.T) {
	// We are declaring a slice of structs by using []struct with two fields, the shape and want
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// Test cases
		// We can rename the fields here to provide some descriptive test cases so
		// in case if one fails we will be able to pinpoint the issue.
		// {Rectangle{12, 6}, 72.0},
		// {Circles{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circles{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

// Above code can further be refactored to below:
func TestAreaV7(t *testing.T) {
	// We are declaring a slice of structs by using []struct with two fields, the shape and want
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		// Test cases
		// We can rename the fields here to provide some descriptive test cases so
		// in case if one fails we will be able to pinpoint the issue.
		// {Rectangle{12, 6}, 72.0},
		// {Circles{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.10},
		{name: "Circles", shape: Circles{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.10},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
