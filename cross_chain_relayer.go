package main

import (
	"fmt"
	"time"
)

type Chain struct {
	ChainID string
	Height  int
}

type CrossChainRelayer struct {
	SourceChain *Chain
	TargetChain *Chain
}

func NewRelayer(source, target *Chain) *CrossChainRelayer {
	return &CrossChainRelayer{
		SourceChain: source,
		TargetChain: target,
	}
}

func (r *CrossChainRelayer) RelayData(data string) string {
	fmt.Printf("Relaying data from %s to %s...\n", r.SourceChain.ChainID, r.TargetChain.ChainID)
	time.Sleep(200 * time.Millisecond)
	proof := fmt.Sprintf("PROOF_%s_%d", data, time.Now().Unix())
	return proof
}

func (r *CrossChainRelayer) VerifyRelay(proof string) bool {
	return len(proof) > 10
}

func main() {
	chainA := &Chain{ChainID: "NEXUS", Height: 2500}
	chainB := &Chain{ChainID: "ETHEREUM", Height: 20000000}
	relayer := NewRelayer(chainA, chainB)

	proof := relayer.RelayData("CROSS_CHAIN_TRANSFER_100")
	valid := relayer.VerifyRelay(proof)
	fmt.Printf("Relay Proof: %s...\n", proof[:16])
	fmt.Printf("Relay Verify: %t\n", valid)
}
