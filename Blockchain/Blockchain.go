package Blockchain

import (
	"bytes"
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

func (bc *Blockchain) GenTokenFromFaucet() {

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
	if bytes.Equal(bc.BlockHistory[len(bc.BlockHistory)-1].Sign, b.PrevHash) {
		bc.BlockHistory = append(bc.BlockHistory, b)
	}

}

func (bc *Blockchain) VerifyBlock(b *Block) bool {
	if bytes.Equal(bc.BlockHistory[len(bc.BlockHistory)-1].Sign, b.PrevHash) {
		if b.Verify() {
			return true
		}
	}

	return false
}

func (bc *Blockchain) ShowHistory() {
	fmt.Println("SHOW HISTORY")
	for i, block := range bc.BlockHistory {
		fmt.Println(i)
		fmt.Println(block.ToString())
	}
}

func (bc *Blockchain) ShowMappol() {
	fmt.Println("SHOW MAPPOOL")
	for i, tx := range Mappool {
		fmt.Println(i)
		fmt.Println(tx.ToString())
	}
}
