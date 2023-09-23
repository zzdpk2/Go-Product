package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw Succeeded", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10.0)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	assertError := func(t *testing.T, err error, want error) {
		if err == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if err != want {
			t.Errorf("got '%s', want '%s'", err, want)
		}
	}

	t.Run("Withdraw Failed", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30.0)
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
		assertError(t, err, InsufficientFundsError)
	})

	// t.Run("Withdraw Failed", func(t *testing.T) {
	// 	wallet := Wallet{balance: Bitcoin(20)}
	// 	wallet.Withdraw(30.0)
	// 	want := Bitcoin(20)
	// 	assertBalance(t, wallet, want)
	// })

}
