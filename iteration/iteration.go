package iteration

func Repeatv1(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

// further refactoring
const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}
