package pointers

// Wallet containing a Balance
type Wallet struct {
	balance int
}

// Deposit adds an Amount to the Wallet struct
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns the current Balance
func (w *Wallet) Balance() int {
	return w.balance
}
