package main

import (
	"fmt"
	"sync"
)

type Shard struct {
	ShardID int
	Txs     []string
}

type ShardEngine struct {
	Shards []Shard
}

func (e *ShardEngine) ParallelProcess() {
	var wg sync.WaitGroup
	for i := range e.Shards {
		wg.Add(1)
		go func(s *Shard) {
			defer wg.Done()
			fmt.Printf("Shard %d processing %d txs\n", s.ShardID, len(s.Txs))
		}(&e.Shards[i])
	}
	wg.Wait()
}

func main() {
	engine := ShardEngine{
		Shards: []Shard{
			{0, []string{"tx1", "tx2"}},
			{1, []string{"tx3", "tx4"}},
		},
	}
	engine.ParallelProcess()
}
