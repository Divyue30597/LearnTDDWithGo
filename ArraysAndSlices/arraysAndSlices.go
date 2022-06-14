package ArraysAndSlices

func SumV1(number [6]int) int {
	sum := 0
	for i := 0; i < 6; i++ {
		sum += number[i]
	}
	return sum
}

// Refactoring the code to use range

// range let's us iterate over an array and on each iteration it returns 2 values - the index and the value

func SumV2(numbers [6]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Note 1: An interesting property of arrays is that the size is encoded in its type. If you try to pass an [4]int into a function that expects [5]int, it won't compile. They are different types so it's just the same as trying to pass a string into a function that wants an int.

// Note 2: Size is always encoded in the arrays always but that is not used oftenly. So, to solve this issue Go has slices.

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Variadic Function => can be called with any number of trailing arguments
// The below function has numbersToSum arguments with n number of trailing arguments and will return sums
func SumAllV1(numbersToSum ...[]int) []int {
	// return
	lengthOfNumbers := len(numbersToSum)
	// make -> allows us to create a slice with a starting capacity of the len of the numbersToSum. This is a new way to create slice.
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums

}

// Another way to write the above function is to use append and initialising sums as empty slice
// In below implementation, we are worrying less about capacity. We start with an empty slice sums and append to it the result of Sum as we work through the varargs.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// Problem statement 2:
// change SumAll to SumAllTails, where it will calculate the totals of the "tails" of each slice. The tail of a collection is all items in the collection except the first one (the "head").

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, number := range numbersToSum {
		// Slices can be sliced. 
		tail := number[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
