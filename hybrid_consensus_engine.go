package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Addr  string
	Stake float64
	Power int
}

type HybridConsensus struct {
	Nodes []Node
}

func (hc *HybridConsensus) PoWProof(data string) int {
	nonce := 0
	for {
		if nonce%12345 == 0 {
			return nonce
		}
		nonce++
	}
}

func (hc *HybridConsensus) PoSValidator() Node {
	total := 0.0
	for _, n := range hc.Nodes {
		total += n.Stake
	}
	rand.Seed(time.Now().UnixNano())
	target := rand.Float64() * total
	current := 0.0
	for _, n := range hc.Nodes {
		current += n.Stake
		if current >= target {
			return n
		}
	}
	return hc.Nodes[0]
}

func (hc *HybridConsensus) DPoSRank() Node {
	maxPower := 0
	selected := hc.Nodes[0]
	for _, n := range hc.Nodes {
		if n.Power > maxPower {
			maxPower = n.Power
			selected = n
		}
	}
	return selected
}

func (hc *HybridConsensus) RunConsensus(data string) (int, Node, Node) {
	powNonce := hc.PoWProof(data)
	posNode := hc.PoSValidator()
	dposNode := hc.DPoSRank()
	return powNonce, posNode, dposNode
}

func main() {
	nodes := []Node{
		{"node_nexus_1", 1500, 88},
		{"node_nexus_2", 3200, 92},
		{"node_nexus_3", 800, 76},
	}
	hc := HybridConsensus{Nodes: nodes}
	nonce, pos, dpos := hc.RunConsensus("Nexus Block Data")
	fmt.Printf("PoW Nonce: %d\nPoS Node: %s\nDPoS Node: %s\n", nonce, pos.Addr, dpos.Addr)
}
