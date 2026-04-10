package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateAnonymousAddress() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return "ANON_" + hex.EncodeToString(buf)
}

func main() {
	addr := GenerateAnonymousAddress()
	fmt.Printf("Anonymous Address: %s\n", addr)
}
