package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Addr      string
	StakedAmt float64
	IsActive  bool
}

type StakingCore struct {
	Validators []Validator
}

func (s *StakingCore) ElectValidator() Validator {
	rand.Seed(time.Now().UnixNano())
	return s.Validators[rand.Intn(len(s.Validators))]
}

func (s *StakingCore) DistributeReward(validator *Validator, reward float64) {
	validator.StakedAmt += reward
	fmt.Printf("Reward Distributed! New Stake: %.2f\n", validator.StakedAmt)
}

func main() {
	core := &StakingCore{
		Validators: []Validator{
			{"val1", 5000, true},
			{"val2", 8000, true},
			{"val3", 3500, true},
		},
	}
	selected := core.ElectValidator()
	core.DistributeReward(&selected, 100.5)
	fmt.Printf("Elected Validator: %s\n", selected.Addr)
}
