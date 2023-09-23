package main

import "fmt"

// type Bitcoin float64

type Wallet struct {
	balance float64
}

func (w *Wallet) Balance() float64 {
	fmt.Println("address of balance in Deposit is", &w.balance)
	return w.balance
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}
