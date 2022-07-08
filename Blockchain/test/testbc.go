package main

import (
	"Lab/BabyBlockChain2/Blockchain"
	"fmt"
)

func main() {
	ac := Blockchain.GenAccount()
	ac.ChangeMyStatus()
	ac2 := Blockchain.GenAccount()
	ac.BecomeCandidate(ac2)
	op, _ := ac.CreateOperation(*ac2)
	tx := ac.CreateTxt()
	tx.AddOp(*op)
	bc := Blockchain.InitBlockchain()

	b1 := ac.CreateBlock(bc.BlockHistory[len(bc.BlockHistory)-1].Sign)
	b1.AddTx()
	bc.AddBlock(b1)
	b2 := ac.CreateBlock(bc.BlockHistory[len(bc.BlockHistory)-1].Sign)
	b1.AddTx()
	bc.AddBlock(b2)

	fmt.Println(b1.ToString())

	bc.ShowCoindata()

}
