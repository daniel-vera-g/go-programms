package pointers

type Bitcoin Bitcoin

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
