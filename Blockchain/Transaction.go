package Blockchain

import (
	"crypto/rsa"
	"fmt"
)

var (
	ID = 0
)

var Mappool = make([]Transaction, 0)

type Transaction struct {
	ID        int
	Operation *Operation
	nonce     uint
	FullSign  []byte
	PublicKey *rsa.PublicKey
	numOfSing int8
}

func (tx *Transaction) signTX(a *Account) ([]byte, error) {

	sign, err := a.SignData(tx.ToString())
	tx.FullSign = sign
	tx.PublicKey = a.Wallets.PublicKey

	return sign, err
}

func (tx *Transaction) AddOp(o Operation) {

	if &o != nil {
		if VerifyOperation(o) {
			tx.Operation = &o
			tx.FullSign, _ = tx.signTX(&o.Sender)
			tx.PublicKey = tx.Operation.Sender.Wallets.PublicKey
			tx.addToMappool()
		}

	}
}

func (tx *Transaction) addToMappool() {
	Mappool = append(Mappool, *tx)
}

func (tx *Transaction) ToString() string {
	return fmt.Sprintf("ID - %d \n Op - %s\n Nonce - %d\n", tx.ID, tx.Operation.ToString(), tx.nonce)
}
