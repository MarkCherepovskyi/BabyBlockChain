package main

import (
	"Lab/BabyBlockChain2/Account"
	"crypto/rsa"
	"fmt"
)

func check(key *rsa.PublicKey, msg string, sign []byte) {
	if Account.Verify(key, msg, sign) {
		fmt.Println("Nice")
	} else {
		fmt.Println("Error")
	}
}

func main() {
	//keys := Account.Keys{}
	//
	//keys.PrivateKey, keys.PublicKey = keys.GenKeys()
	//
	//keys.ToString()
	//
	//msg := "msg"
	//msg2 := "MSG222"
	//sign := keys.Sign(msg, keys.PrivateKey)
	//
	//check(keys.PublicKey, msg, sign)
	//check(keys.PublicKey, msg2, sign)
	user := Account.GenAccount()
	user.ToString()
	msg := "msg 1"
	sign, _ := user.SignData(msg)
	if Account.Verify(user.Wallets.PublicKey, msg, sign) {
		fmt.Println("Nice")
	}

	user2 := Account.GenAccount()

	op1, err := user.CreateOperation(*user2, 2)
	if err != nil {
		return
	}
	op2, err := user.CreateOperation(*user2, 5)
	if err != nil {
		return
	}
	op3, err := user2.CreateOperation(*user, 5)
	if err != nil {
		return
	}
	tx := user.CreateTxt()
	tx.AddTX(*op1)
	tx.AddTX(*op2)
	tx.AddTX(*op3)

	fmt.Println("//////////////")
	fmt.Println(tx.ToString())

}
