package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

type QuantumCrypto struct{}

func (qc *QuantumCrypto) GenerateQuantumKey() (string, string) {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	pub := sha256.Sum256(key)
	return hex.EncodeToString(key), hex.EncodeToString(pub[:])
}

func (qc *QuantumCrypto) QuantumSign(privateKey, message string) string {
	combined := privateKey + message
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

func (qc *QuantumCrypto) QuantumVerify(publicKey, message, signature string) bool {
	verifySign := qc.QuantumSign(publicKey, message)
	return verifySign == signature
}

func main() {
	qc := QuantumCrypto{}
	priv, pub := qc.GenerateQuantumKey()
	msg := "Nexus Quantum Transaction"
	sign := qc.QuantumSign(priv, msg)
	valid := qc.QuantumVerify(pub, msg, sign)

	fmt.Printf("Private Key: %s...\n", priv[:16])
	fmt.Printf("Public Key: %s...\n", pub[:16])
	fmt.Printf("Signature: %s...\n", sign[:16])
	fmt.Printf("Quantum Verify: %t\n", valid)
}
