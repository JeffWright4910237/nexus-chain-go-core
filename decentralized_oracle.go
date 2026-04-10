package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DecentralizedOracle struct {
	Sources []string
}

func (o *DecentralizedOracle) FetchData() float64 {
	rand.Seed(time.Now().UnixNano())
	return 18000 + rand.Float64()*1000
}

func (o *DecentralizedOracle) VerifyData(data float64) bool {
	return data > 17000 && data < 19000
}

func main() {
	oracle := DecentralizedOracle{Sources: []string{"Binance", "CoinGecko", "Chainlink"}}
	price := oracle.FetchData()
	valid := oracle.VerifyData(price)
	fmt.Printf("Oracle Price: %.2f\n", price)
	fmt.Printf("Data Valid: %t\n", valid)
}
