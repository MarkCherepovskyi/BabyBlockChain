package main

import (
	"Lab/BabyBlockChain2/Account"
	"crypto/rsa"
	"fmt"
)

func check(key *rsa.PublicKey, msg string, sign []byte) {
	buf, _ := Account.Verify(key, msg, sign)
	if buf {
		fmt.Println("Nice")
	} else {
		fmt.Println("Error")
	}
}

func main() {

	user := Account.GenAccount()
	user.ToString()
	msg := "msg 1"
	sign, _ := user.SignData(msg)
	buf, _ := Account.Verify(user.Wallets.PublicKey, msg, sign)
	if buf {
		fmt.Println("Nice")
	}

	user2 := Account.GenAccount()

	op1, err := user.CreateOperation(*user2, 2)
	if err != nil {
		return
	}

	tx := user.CreateTxt()
	tx.AddTX(*op1)

	fmt.Println("//////////////")
	fmt.Println(tx.ToString())

	user.ChangeMyStatus()
	user.SignTX(tx)
	fmt.Println(tx.FullSign)
	fmt.Println(tx.PublicKey)
	if user.VerifyTX(tx) {
		fmt.Println("Nice. TX is valid")
	} else {
		fmt.Println("All is bad ")
	}

	user.ShowMappol()

}
