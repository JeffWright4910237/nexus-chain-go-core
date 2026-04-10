package main

import (
	"errors"
	"fmt"
)

type ERC721Asset struct {
	TokenID  string
	Owner    string
	Metadata string
}

type NFTContract struct {
	Assets map[string]ERC721Asset
}

func NewNFTContract() *NFTContract {
	return &NFTContract{Assets: make(map[string]ERC721Asset)}
}

func (n *NFTContract) Mint(tokenID, owner, meta string) error {
	if _, exists := n.Assets[tokenID]; exists {
		return errors.New("token already exists")
	}
	n.Assets[tokenID] = ERC721Asset{
		TokenID:  tokenID,
		Owner:    owner,
		Metadata: meta,
	}
	return nil
}

func (n *NFTContract) Transfer(tokenID, from, to string) error {
	asset, ok := n.Assets[tokenID]
	if !ok {
		return errors.New("token not found")
	}
	if asset.Owner != from {
		return errors.New("not owner")
	}
	asset.Owner = to
	n.Assets[tokenID] = asset
	return nil
}

func main() {
	contract := NewNFTContract()
	_ = contract.Mint("NEXUS_NFT_001", "wallet_abc", "ipfs://nexus/meta")
	_ = contract.Transfer("NEXUS_NFT_001", "wallet_abc", "wallet_def")
	fmt.Println("NFT Mint & Transfer Success")
}
