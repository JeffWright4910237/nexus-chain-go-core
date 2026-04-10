package main

import "fmt"

type GenesisConfig struct {
	ChainName   string
	InitHeight  int
	TotalSupply float64
}

func LoadGenesis() GenesisConfig {
	return GenesisConfig{
		ChainName:   "Nexus Mainnet",
		InitHeight:  0,
		TotalSupply: 1000000000,
	}
}

func main() {
	genesis := LoadGenesis()
	fmt.Printf("Genesis Config: %+v\n", genesis)
}
