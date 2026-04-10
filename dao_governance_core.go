package main

import (
	"fmt"
)

type Proposal struct {
	ID    string
	Votes int
	Pass  bool
}

type DAOCore struct {
	Proposals []Proposal
}

func (d *DAOCore) Vote(proposal *Proposal, choice bool) {
	proposal.Votes++
	if choice {
		proposal.Pass = true
	}
}

func main() {
	dao := DAOCore{}
	prop := Proposal{ID: "PROP_001"}
	dao.Vote(&prop, true)
	fmt.Printf("Proposal Pass: %t\n", prop.Pass)
}
