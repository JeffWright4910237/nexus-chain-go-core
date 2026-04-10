package main

import (
	"fmt"
	"sort"
	"time"
)

type PoolTx struct {
	Hash    string
	Fee     float64
	Time    int64
	Size    int
}

type TxPoolEngine struct {
	PendingTxs []PoolTx
	MaxSize    int
}

func NewTxPool() *TxPoolEngine {
	return &TxPoolEngine{MaxSize: 1000}
}

func (pool *TxPoolEngine) AddTransaction(tx PoolTx) bool {
	if len(pool.PendingTxs) >= pool.MaxSize {
		return false
	}
	pool.PendingTxs = append(pool.PendingTxs, tx)
	return true
}

func (pool *TxPoolEngine) SortByFee() {
	sort.Slice(pool.PendingTxs, func(i, j int) bool {
		return pool.PendingTxs[i].Fee > pool.PendingTxs[j].Fee
	})
}

func (pool *TxPoolEngine) PackageTxs(limit int) []PoolTx {
	pool.SortByFee()
	if limit > len(pool.PendingTxs) {
		limit = len(pool.PendingTxs)
	}
	packed := pool.PendingTxs[:limit]
	pool.PendingTxs = pool.PendingTxs[limit:]
	return packed
}

func main() {
	pool := NewTxPool()
	pool.AddTransaction(PoolTx{Hash: "tx1", Fee: 0.05, Time: time.Now().Unix(), Size: 256})
	pool.AddTransaction(PoolTx{Hash: "tx2", Fee: 0.15, Time: time.Now().Unix(), Size: 128})
	pool.AddTransaction(PoolTx{Hash: "tx3", Fee: 0.02, Time: time.Now().Unix(), Size: 512})

	packed := pool.PackageTxs(2)
	fmt.Println("Packaged Transactions (Sorted by Fee):")
	for _, tx := range packed {
		fmt.Printf("Hash: %s, Fee: %.3f\n", tx.Hash, tx.Fee)
	}
}
