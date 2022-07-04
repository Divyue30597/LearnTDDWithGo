package pointeranderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	// When a method or a function is called, the arguments are copied.
	// When calling func (w Wallet) Deposit(amount int) the w is a copy of
	// whatever we called the method from.
	wallet.Deposit(Bitcoin(1))

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %v %v \n", &wallet.balance, wallet.balance)

	// To make Bitcoin you just use the syntax Bitcoin(999).
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

/*
  OUTPUT :
	Address of balance in Deposit is 0xc0000142c8 0
	address of balance in test is 0xc0000142c8 9

	You can see that the addresses of the two balances are different. So when we change the value of the balance inside the code, we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.

	To Fix this we can use Pointers -> let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.
*/

func TestWallet2(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		fmt.Printf("address of balance in test is %v %v \n", &wallet.balance, wallet.balance)

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdrawal(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

// Refactoring the above code.
func TestWallet3(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdrawal(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawal(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("Wanted an error but didnot get one.")
		}
	})
}

// Further refactor
func TestWallet4(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but didnot get one.")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdrawal(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawal(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}

// Further refactoring taking the functions out
func TestWallet5(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdrawal(Bitcoin(10))
		assetNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawal(Bitcoin(100))

		// The var keyword allows us to define values global to the package.
		// This is a positive change in itself because now our Withdraw
		// function looks very clear.
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didnot get one.")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assetNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but never wanted one")
	}
}

/* 
	Summary - 
	Pointers
	Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
	The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

	nil
	Pointers can be nil.
	When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
	Useful for when you want to describe a value that could be missing.

	Errors
	Errors are the way to signify failure when calling a function/method.
	By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
	This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
*/