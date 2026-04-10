package main

import (
	"errors"
	"fmt"
)

type MultiSigWallet struct {
	Owners   []string
	Required int
}

func (w *MultiSigWallet) SignTx(signers []string) error {
	if len(signers) < w.Required {
		return errors.New("insufficient signatures")
	}
	return nil
}

func main() {
	wallet := MultiSigWallet{
		Owners:   []string{"owner1", "owner2", "owner3"},
		Required: 2,
	}
	err := wallet.SignTx([]string{"owner1", "owner2"})
	if err == nil {
		fmt.Println("MultiSig Transaction Signed ✅")
	}
}
