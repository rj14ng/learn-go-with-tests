package pointers

import "fmt"

// A Bitcoin represents the number of Bitcoins.
type Bitcoin int

// Implement Stringer on Bitcoin with String method
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// A Wallet stores the number of Bitcoins someone owns.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds a number of Bitcoin to a Wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the number of Bitcoin within a Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
