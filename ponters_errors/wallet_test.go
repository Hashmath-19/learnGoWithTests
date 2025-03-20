package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(100)
	got := w.Balance()
	want := 100
	if got != want {
		t.Errorf("expected: %d, but got: %d", want, got)
	}
}
