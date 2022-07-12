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
		[]byte("GENESIS_BLOCK"),
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
	if bytes.Equal(bc.LastHash(), b.PrevHash) && b.setOfTransaction != nil {
		bc.BlockHistory = append(bc.BlockHistory, b)
		bc.AddToTxDB(b)
	}

}

func (bc *Blockchain) AddToTxDB(b *Block) {
	for _, tx := range b.setOfTransaction {
		bc.TxDatabase = append(bc.TxDatabase, tx)
	}
}

func (bc *Blockchain) VerifyBlock(b *Block) bool {
	if bytes.Equal(bc.LastHash(), b.PrevHash) {
		if b.Verify() {
			return true
		}
	}

	return false
}

func (bc *Blockchain) ShowHistory() {
	fmt.Println("SHOW HISTORY")
	for _, block := range bc.BlockHistory {

		fmt.Println(block.ToString())
	}
}

func (bc *Blockchain) ShowMappol() {
	fmt.Println("SHOW MAPPOOL")
	for _, tx := range Mappool {

		fmt.Println(tx.ToString())
	}
}
func (bc *Blockchain) ShowScore() {
	fmt.Println("SHOW SCORE")
	for _, candidate := range bc.Candidates {
		fmt.Printf("Candidate %s has  %d coins", candidate.ID, candidate.Balance)
	}
}

func (bc *Blockchain) ShowCandidates() {
	fmt.Println("SHOW CANDIDATES")
	for _, candidate := range bc.Candidates {
		fmt.Printf("Candidate %s ", candidate.ID)
	}
}

func (bc *Blockchain) ShowLen() {
	fmt.Printf("Len is %d blocks", len(bc.BlockHistory))
}

func (bc *Blockchain) ShowTXDB() {
	fmt.Println("SHOW TX DB")
	for _, tx := range bc.TxDatabase {
		fmt.Println(tx)
	}
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.BlockHistory[len(bc.BlockHistory)-1]
}

func (bc *Blockchain) LastHash() []byte {
	hash := bc.LastBlock().Sign
	if hash == nil {
		return []byte("Error hash ")
	}
	return hash
}
