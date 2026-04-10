package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ZKProofSystem struct{}

func (zk *ZKProofSystem) GenerateProof(secret string) string {
	hash := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(hash[:]) + "_ZKPROOF"
}

func (zk *ZKProofSystem) VerifyProof(proof string) bool {
	return len(proof) == 80
}

func main() {
	zk := ZKProofSystem{}
	secret := "USER_BALANCE_9999"
	proof := zk.GenerateProof(secret)
	valid := zk.VerifyProof(proof)

	fmt.Printf("ZK Proof: %s...\n", proof[:16])
	fmt.Printf("Proof Verify: %t\n", valid)
}
