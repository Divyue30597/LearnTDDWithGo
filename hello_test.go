package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	got := Hello2("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Example of subtests. Here we are introducing subtests.
// Sometimes it is useful to group tests around a thing and then have subtests describing different scenarios
func TestHello3(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello2("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello3("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

// The above code can further be refactored to the code below

func TestHello4(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello2("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello3("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func TestHello5(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello4("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello4("Divyue", "French")
		want := "Bonjour, Divyue"
		assertCorrectMessage(t, got, want)
	})
}
