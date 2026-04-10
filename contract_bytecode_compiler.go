package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type ContractCompiler struct{}

func (c *ContractCompiler) Compile(source string) (string, string) {
	bytecode := strings.ToUpper(hex.EncodeToString([]byte(source)))
	abiHash := sha256.Sum256([]byte(bytecode))
	abi := hex.EncodeToString(abiHash[:])
	return bytecode, abi
}

func main() {
	compiler := ContractCompiler{}
	source := "function transfer() { return true; }"
	bytecode, abi := compiler.Compile(source)
	fmt.Printf("Bytecode: %s...\n", bytecode[:16])
	fmt.Printf("ABI Hash: %s...\n", abi[:16])
}
