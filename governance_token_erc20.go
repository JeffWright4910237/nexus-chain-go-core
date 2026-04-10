package main

import (
	"errors"
	"fmt"
)

type ERC20Token struct {
	Name     string
	Symbol   string
	Decimals int
	Total    float64
	Balances map[string]float64
}

func NewToken() *ERC20Token {
	return &ERC20Token{
		Name:     "Nexus Governance Token",
		Symbol:   "NEX",
		Decimals: 18,
		Total:    100000000,
		Balances: make(map[string]float64),
	}
}

func (t *ERC20Token) Transfer(from, to string, amount float64) error {
	if t.Balances[from] < amount {
		return errors.New("insufficient balance")
	}
	t.Balances[from] -= amount
	t.Balances[to] += amount
	return nil
}

func main() {
	token := NewToken()
	token.Balances["wallet_a"] = 10000
	_ = token.Transfer("wallet_a", "wallet_b", 500)
	fmt.Printf("Wallet B Balance: %.2f\n", token.Balances["wallet_b"])
}
