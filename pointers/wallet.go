package pointers

import (
	"errors"
	"fmt"
)

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

// ErrInsufficientFunds means a Wallet does not have enough Bitcoin to perform a Withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw removes a number of Bitcoin from a Wallet, returning an error if it cannot be performed.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	// Insufficient funds
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the number of Bitcoin within a Wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
