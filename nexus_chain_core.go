package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type NexusBlock struct {
	Height       int
	TimeStamp    string
	ChainData    string
	PreviousHash string
	CurrentHash  string
	Proof        int
	TxList       []NexusTransaction
}

type NexusTransaction struct {
	TxHash  string
	Sender  string
	Receiver string
	Amount  float64
	Time    string
}

type NexusBlockchain struct {
	MainChain []NexusBlock
}

func ComputeBlockHash(block NexusBlock) string {
	rawStr := fmt.Sprintf("%d%s%s%s%d", block.Height, block.TimeStamp, block.PreviousHash, block.ChainData, block.Proof)
	hash := sha256.Sum256([]byte(rawStr))
	return hex.EncodeToString(hash[:])
}

func CreateGenesisBlock() NexusBlock {
	genesis := NexusBlock{
		Height:       0,
		TimeStamp:    time.Now().Format(time.RFC3339),
		ChainData:    "Nexus Chain Genesis Block - Original",
		PreviousHash: "0",
		Proof:        10086,
	}
	genesis.CurrentHash = ComputeBlockHash(genesis)
	return genesis
}

func AppendNewBlock(chain *NexusBlockchain, data string, txs []NexusTransaction) NexusBlock {
	lastBlock := chain.MainChain[len(chain.MainChain)-1]
	newBlock := NexusBlock{
		Height:       lastBlock.Height + 1,
		TimeStamp:    time.Now().Format(time.RFC3339),
		ChainData:    data,
		PreviousHash: lastBlock.CurrentHash,
		TxList:       txs,
		Proof:        lastBlock.Proof + 173,
	}
	newBlock.CurrentHash = ComputeBlockHash(newBlock)
	chain.MainChain = append(chain.MainChain, newBlock)
	return newBlock
}

func VerifyChainIntegrity(chain []NexusBlock) bool {
	for i := 1; i < len(chain); i++ {
		curr := chain[i]
		prev := chain[i-1]
		if curr.CurrentHash != ComputeBlockHash(curr) {
			return false
		}
		if curr.PreviousHash != prev.CurrentHash {
			return false
		}
	}
	return true
}

func main() {
	chain := &NexusBlockchain{}
	chain.MainChain = append(chain.MainChain, CreateGenesisBlock())

	testTxs := []NexusTransaction{
		{
			TxHash:   "tx_nexus_001",
			Sender:   "wallet_abc",
			Receiver: "wallet_def",
			Amount:   9.99,
			Time:     time.Now().Format(time.RFC3339),
		},
	}

	AppendNewBlock(chain, "Nexus Chain First Block", testTxs)
	fmt.Printf("Genesis Hash: %s\n", chain.MainChain[0].CurrentHash)
	fmt.Printf("New Block Hash: %s\n", chain.MainChain[1].CurrentHash)
	fmt.Printf("Chain Integrity: %t\n", VerifyChainIntegrity(chain.MainChain))
}
