package main

import (
	"Lab/BabyBlockChain2/Blockchain"
	"log"
)

func main() {

	bc := Blockchain.InitBlockchain()

	ac := Blockchain.GenAccount()
	ac.ChangeMyStatus()
	ac2 := Blockchain.GenAccount()
	ac.BecomeCandidate(ac2, bc)
	ac3 := Blockchain.GenAccount()
	ac4 := Blockchain.GenAccount()
	ac5 := Blockchain.GenAccount()

	_, err := ac.CreateOperation(ac2)
	if err != nil {
		log.Println(err)
	}
	_, err = ac3.CreateOperation(ac2)
	if err != nil {
		log.Println(err)
	}
	_, err = ac3.CreateOperation(ac2)
	if err != nil {
		log.Println(err)
	}
	_, err = ac4.CreateOperation(ac2)
	if err != nil {
		log.Println(err)
	}
	_, err = ac5.CreateOperation(ac2)
	if err != nil {
		log.Println(err)
	}

	ac.CreateBlock(bc)
	bc.ShowLen()

	ac.CreateBlock(bc)
	bc.ShowLen()

	ac.CreateBlock(bc)
	bc.ShowLen()
	//
	//fmt.Println(b1.ToString())

	bc.ShowMappol()
	bc.ShowHistory()

	bc.ShowCandidates()
	bc.ShowScore()
	bc.ShowLen()
	
}
