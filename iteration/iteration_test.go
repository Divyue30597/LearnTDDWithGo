package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

// Benchmarking

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestExampleRepeat(t *testing.T) {
	repeated := ExampleRepeat("a", 7)
	expected := "aaaaaaa"
	if repeated != expected {
		t.Errorf("Expected %q got %q", expected, repeated)
	}
}
