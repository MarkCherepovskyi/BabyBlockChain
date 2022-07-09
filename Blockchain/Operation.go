package Blockchain

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Operation struct {
	Sender    Account
	Receiver  Account
	Amount    int
	Signature []byte
}

func VerifyOperation(o Operation) bool {
	if o.Amount > o.Sender.Balance {
		fmt.Println("Balance is unvalid")
		return false
	}
	buf, _ := Verify(o.Sender.Wallets.GetPublicKey(), o.ToString(), o.Signature)
	if !buf {
		fmt.Println("Someone changed the data")
		return false
	}
	return true
}

func (o *Operation) CreateTxt() *Transaction {
	randNum, _ := rand.Int(rand.Reader, big.NewInt(100000))

	tx := Transaction{
		ID,
		nil,
		randNum.Uint64(),
		nil,
		nil,
		0,
	}
	ID++
	tx.addOp(*o)

	return &tx
}

func (o Operation) ToString() string {
	senderStr := o.Sender.ToString()
	receiverStr := o.Receiver.ToString()

	return fmt.Sprintf("Sender: \n%s\nReceiver: \n%s\nAmount -  %d\n", senderStr, receiverStr, o.Amount)
}

func (o Operation) PrintKey() {
	fmt.Println(o.Sender.Wallets.ToString())
}

func (o Operation) carryOutOp() {
	o.Receiver.Balance++
	o.Sender.Balance--

}
