package main

import (
	"fmt"
)

// Ops 是基本的货币接口，定义了所有货币类型都需要实现的方法。
type Ops interface {
	GetAmount() float64
	SetAmount(amount float64)
}

type Crypto struct {
	amount float64
	unit   string
}

// Ethereum 是 Crypto 的一个实现。
type Ethereum struct {
	Crypto
}

func (e *Ethereum) GetAmount() float64 {
	return e.amount
}

func (e *Ethereum) SetAmount(amount float64) {
	e.amount = amount
}

func (e Ethereum) String() string {
	return fmt.Sprintf("%f ETH", e.amount)
}

// Bitcoin 也是 Crypto 的一个实现。
type Bitcoin struct {
	Crypto
}

func (b *Bitcoin) GetAmount() float64 {
	return b.amount
}

func (b *Bitcoin) SetAmount(amount float64) {
	b.amount = amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b.amount)
}

type Wallet struct {
	currency Crypto
}

func (w *Wallet) Deposit(amount Crypto) {
	switch w.currency.unit {
	case "ETH":
		if amount.unit == "ETH" {
			w.currency.amount += amount.amount
		} else {
			fmt.Println("Mismatched currency type: expected Ethereum")
		}
	case "BTC":
		if amount.unit == "BTC" {
			w.currency.amount += amount.amount
		} else {
			fmt.Println("Mismatched currency type: expected Bitcoin")
		}
	default:
		fmt.Println("Unknown currency type in the wallet")
	}
}

func (w *Wallet) Withdraw(amount Crypto) {
	currentAmount := w.currency.amount
	if currentAmount < amount.amount {
		fmt.Println("Not enough balance!")
		return
	}
	// w.currency.SetAmount(currentAmount - amount.amount)
	w.currency.amount = currentAmount - amount.amount
}

func (w *Wallet) Balance() float64 {
	return w.currency.amount
}

func NewWallet(currency Crypto) *Wallet {
	return &Wallet{currency: currency}
}
