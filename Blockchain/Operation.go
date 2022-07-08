package Blockchain

import (
	"fmt"
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
	buf, _ := Verify(o.Sender.Wallets.PublicKey, o.ToString(), o.Signature)
	if buf {
		fmt.Println("Has anyone changed the data")
		return false
	}
	return true
}

func (o Operation) ToString() string {
	senderStr := o.Sender.ToString()
	receiverStr := o.Receiver.ToString()
	return fmt.Sprintf("%s %s %d", senderStr, receiverStr, o.Amount)
}

func (o Operation) PrintKey() {
	fmt.Println(o.Sender.Wallets.ToString())
}

func (o Operation) carryOutOp() {
	o.Receiver.Balance++
	o.Sender.Balance--

}
