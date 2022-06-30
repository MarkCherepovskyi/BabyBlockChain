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

}
