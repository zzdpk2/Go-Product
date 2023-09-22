package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Crypto{
		amount: 10,
		unit:   "ETH",
	})
	got := wallet.Balance()
	fmt.Println("address of balance in test is", wallet.Balance())
	want := Ethereum{
		Crypto{
			amount: 10,
			unit:   "ETH",
		}}
	fmt.Printf("got: %f", got)
	fmt.Printf("want: %s", want)

	want1 := want.amount
	epsilon := Ethereum{
		Crypto{
			amount: 0.000001,
			unit:   "ETH",
		}}
	if (got - want1) > epsilon.GetAmount() {
		t.Errorf("got %f want %f", got, want1)
	}
}
