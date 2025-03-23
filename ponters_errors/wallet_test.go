package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(100)
		assertBalance(t, w, BitCoin(100))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{balance: BitCoin(100)}
		// Running errcheck in current dir helps identify 
		// where we have not checked the error being returned on that line of code
		// example: calling w.Withdraw(BitCoin(54)) instead of err := w.Withdraw(BitCoin(54))
		err := w.Withdraw(BitCoin(54))
		assertNoError(t, err)
		assertBalance(t, w, BitCoin(46))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(19)
		w := Wallet{balance: startingBalance}
		err := w.Withdraw(21)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, startingBalance)
	})
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Wanted error but didn't get one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected nil but got error")
	}
}
