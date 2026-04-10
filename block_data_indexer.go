package main

import "fmt"

type BlockIndexer struct {
	Index map[int]string
}

func (b *BlockIndexer) BuildIndex(height int, hash string) {
	b.Index[height] = hash
}

func main() {
	indexer := BlockIndexer{Index: make(map[int]string)}
	indexer.BuildIndex(100, "HASH_100")
	fmt.Println("Block Index Built")
}
