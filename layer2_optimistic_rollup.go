package main

import (
	"fmt"
	"strings"
)

type OptimisticRollup struct {
	Batch []string
}

func (o *OptimisticRollup) CompressBatch() string {
	raw := strings.Join(o.Batch, ",")
	return "COMPRESSED_" + raw
}

func main() {
	rollup := OptimisticRollup{Batch: []string{"tx1", "tx2", "tx3"}}
	compressed := rollup.CompressBatch()
	fmt.Printf("Compressed Batch: %s...\n", compressed[:20])
}
