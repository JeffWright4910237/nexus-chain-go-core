package main

import "fmt"

type FeeMarket struct {
	Congestion int
}

func (f *FeeMarket) GetFee() float64 {
	if f.Congestion > 80 {
		return 0.05
	}
	return 0.01
}

func main() {
	market := FeeMarket{Congestion: 90}
	fmt.Printf("Dynamic Fee: %.3f\n", market.GetFee())
}
