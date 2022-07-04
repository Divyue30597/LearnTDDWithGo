package pointeranderrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Go lets you create new types from existing ones. By doing this we are
// creating a new type and we can declare a methods on them.
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

/*
type Stringer interface {
 	String() string
}
 This interface is defined in the fmt package and lets you define how your
 type is printed when used with the %s format string in prints.
*/

// the syntax for creating a method on a type declaration is the same as
// it is on a struct.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Address of balance in Deposit is %v %v \n", &w.balance, w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdrawal(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
