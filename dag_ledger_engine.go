package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type DAGLedger struct {
	Nodes map[string]DAGNode
}

type DAGNode struct {
	Hash  string
	Prev  []string
	Tx    string
}

func (d *DAGLedger) AddDAGNode(tx string, prev []string) string {
	hash := sha256.Sum256([]byte(tx))
	nodeHash := hex.EncodeToString(hash[:])
	d.Nodes[nodeHash] = DAGNode{Hash: nodeHash, Prev: prev, Tx: tx}
	return nodeHash
}

func main() {
	ledger := DAGLedger{Nodes: make(map[string]DAGNode)}
	node := ledger.AddDAGNode("DAG_TX_001", []string{"genesis"})
	fmt.Printf("DAG Node Hash: %s...\n", node[:16])
}
