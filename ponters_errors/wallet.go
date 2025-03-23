package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go, when you call a function or a method the arguments are copied
// These pointers to structs even have their own name: struct pointers
// and they are automatically dereferenced.
func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount BitCoin) error{
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}
