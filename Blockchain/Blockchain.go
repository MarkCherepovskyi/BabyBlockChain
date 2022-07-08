package Blockchain

import (
	"fmt"
	"time"
)

const (
	GENESIS_BLOCK = "GENESIS_BLOCK"
)

type Blockchain struct {
	CoinDatabase map[string]int
	BlockHistory []*Block
	TxDatabase   []*Transaction
	Candidates   []*Account
	//faucetCoins  int
}

func genesiBlock() *Block {
	b := Block{

		GENESIS_BLOCK,
		nil,
		time.Now(),
		nil,
		nil,
		nil,
	}
	return &b
}

func InitBlockchain() *Blockchain {
	bc := Blockchain{
		nil,
		nil,
		nil,
		nil,
	}

	bc.BlockHistory = append(bc.BlockHistory, genesiBlock())
	return &bc
}

func (bc *Blockchain) AddCandidates(a *Account) {
	if a.Candidate {
		bc.Candidates = append(bc.Candidates, a)
	} else {
		fmt.Println("This account isn't candidate")
	}
}

func (bc *Blockchain) ShowCandidate() {
	fmt.Println("SHOW ALL CANDIDATES")
	for _, data := range bc.Candidates {
		fmt.Println(data.ToString())
	}
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.BlockHistory = append(bc.BlockHistory, b)
}

func (bc *Blockchain) VerifyBlock(b *Block) bool {
	if b.Verify() {
		return true
	}
	return false
}

func (bc *Blockchain) ShowCoindata() {
	fmt.Println("SHOW COINDATABASE")
	for id, balance := range bc.CoinDatabase {
		fmt.Printf("ID - %s\n BALANCE - %d", id, balance)
	}
}
