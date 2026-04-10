package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func OfflineSign(tx string, key string) string {
	sig := sha256.Sum256([]byte(tx + key))
	return hex.EncodeToString(sig[:])
}

func main() {
	sig := OfflineSign("OFFLINE_TX", "SECURE_KEY")
	fmt.Printf("Offline Signature: %s...\n", sig[:16])
}
