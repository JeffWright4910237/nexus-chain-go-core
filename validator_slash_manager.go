package main

import "fmt"

type SlashManager struct {
	PenaltyRate float64
}

func (s *SlashManager) Slash(validator *Validator) {
	validator.StakedAmt *= (1 - s.PenaltyRate)
	fmt.Printf("Slashed! New Stake: %.2f\n", validator.StakedAmt)
}

func main() {
	val := &Validator{StakedAmt: 1000}
	manager := SlashManager{PenaltyRate: 0.1}
	manager.Slash(val)
}
