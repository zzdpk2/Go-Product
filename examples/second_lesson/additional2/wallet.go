package main

import "fmt"

type Bitcoin float64

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Println("address of balance in Deposit is", &w.balance)
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Balance() < amount {
		return fmt.Errorf("insufficient balance %v", amount)
	}

	w.balance -= amount
	return nil
}
