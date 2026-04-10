package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type IPFSStore struct{}

func (i *IPFSStore) Upload(data string) string {
	hash := sha256.Sum256([]byte(data))
	ipfsHash := "Qm" + hex.EncodeToString(hash[:])[:40]
	return ipfsHash
}

func (i *IPFSStore) Verify(hash string) bool {
	return len(hash) == 42
}

func main() {
	store := IPFSStore{}
	hash := store.Upload("Nexus IPFS Storage Data")
	fmt.Printf("IPFS Hash: %s\n", hash)
	fmt.Printf("Verify: %t\n", store.Verify(hash))
}
