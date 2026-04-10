package main

import "fmt"

type WalletConnect struct {
	SessionID string
}

func (w *WalletConnect) Authorize() string {
	return "AUTH_SUCCESS"
}

func main() {
	wallet := WalletConnect{SessionID: "SESS_123"}
	fmt.Println(wallet.Authorize())
}
