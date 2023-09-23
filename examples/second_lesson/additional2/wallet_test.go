package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	// assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
	// 	got := wallet.Balance()
	//
	// 	if got != want {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// }

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw Succeeded", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10.0)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw Failed", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(30.0)

		got := wallet.Balance()

		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
