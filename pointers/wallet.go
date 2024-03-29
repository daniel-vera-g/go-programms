package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet containing a Balance
type Wallet struct {
	balance Bitcoin
}

// Deposit adds an Amount to the Wallet struct
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the current Balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withraw the requested amount from the Wallet
func (w *Wallet) Withraw(amount Bitcoin) error {

	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
