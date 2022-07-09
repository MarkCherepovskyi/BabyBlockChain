package main

import (
	"Lab/BabyBlockChain2/Blockchain"
)

func main() {
	ac := Blockchain.GenAccount()
	ac.ChangeMyStatus()
	ac2 := Blockchain.GenAccount()
	ac.BecomeCandidate(ac2)
	ac3 := Blockchain.GenAccount()
	ac4 := Blockchain.GenAccount()
	ac5 := Blockchain.GenAccount()

	bc := Blockchain.InitBlockchain()

	ac.CreateOperation(*ac2)
	ac3.CreateOperation(*ac2)
	ac4.CreateOperation(*ac2)
	ac5.CreateOperation(*ac2)
	b1 := ac.CreateBlock(bc.BlockHistory[len(bc.BlockHistory)-1].Sign)
	b1.AddTx()
	b2 := ac.CreateBlock(bc.BlockHistory[len(bc.BlockHistory)-1].Sign)
	b2.AddTx()
	bc.AddBlock(b1)
	bc.AddBlock(b2)

	//
	//fmt.Println(b1.ToString())

	bc.ShowMappol()
	bc.ShowHistory()

}
