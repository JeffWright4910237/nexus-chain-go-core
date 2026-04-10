package main

import "fmt"

type LightClient struct {
	SyncHeight int
}

func (l *LightClient) FastSync() {
	l.SyncHeight = 4000
	fmt.Printf("Light Client Synced to Height: %d\n", l.SyncHeight)
}

func main() {
	client := LightClient{}
	client.FastSync()
}
