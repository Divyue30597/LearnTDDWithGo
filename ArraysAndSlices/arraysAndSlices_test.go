// Arrays allows to store multiple elements of the same type in a variable in a particular order.
// Arrays have fixed capacity which can be define when declaring the variable. And Array can be initialized in 2 ways:
// [N]type(value1, value2..., valueN) Eg: [5]int{1,2,3,4,5}
// [...]type(value1, value2..., valueN) Eg: [...]int{1,2,3,4,5}
// %v placeholder is used to print the default format and it works well with arrays

package ArraysAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [6]int{1, 2, 3, 4, 5, 6}

	got := SumV1(numbers)
	want := 21

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumUsingArraysAndSlices(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [6]int{1, 2, 3, 4, 5, 6}

		got := SumV2(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// Having too many tests can turn into a problem and it may just add more overhead in maintenance since Every test has a cost.

// Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.

// SumAll

// Problem statement : For example
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
// or
// SumAll([]int{1,1,1}) would return []int{3}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9, 3})
	want := []int{6, 12}
	// want := "bob"

	// Go does not let us use equality operator with slices.
	// if got != want {
	// 	t.Errorf("Got %v want %v", got, want)
	// }

	// SideEffects: of reflect.DeepEqual is that it is not 'type safe' -> The code will compile if we replace want with "bob"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 3})
	want := []int{5, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
