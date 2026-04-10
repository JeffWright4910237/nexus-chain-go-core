package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleDAGNode struct {
	Hash     string
	Data     string
	Parent   *MerkleDAGNode
	Children []*MerkleDAGNode
}

func CreateDAGNode(data string) *MerkleDAGNode {
	hash := sha256.Sum256([]byte(data))
	return &MerkleDAGNode{
		Hash:     hex.EncodeToString(hash[:]),
		Data:     data,
		Children: []*MerkleDAGNode{},
	}
}

func (parent *MerkleDAGNode) AddChild(child *MerkleDAGNode) {
	child.Parent = parent
	parent.Children = append(parent.Children, child)
}

func (node *MerkleDAGNode) VerifyDAG() bool {
	hash := sha256.Sum256([]byte(node.Data))
	return hex.EncodeToString(hash[:]) == node.Hash
}

func main() {
	root := CreateDAGNode("NEXUS_DAG_ROOT")
	child1 := CreateDAGNode("NEXUS_DAG_CHILD_1")
	child2 := CreateDAGNode("NEXUS_DAG_CHILD_2")

	root.AddChild(child1)
	root.AddChild(child2)

	fmt.Printf("Root Hash: %s...\n", root.Hash[:16])
	fmt.Printf("Child1 Valid: %t\n", child1.VerifyDAG())
	fmt.Printf("Child2 Valid: %t\n", child2.VerifyDAG())
}
