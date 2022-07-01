package Account

import (
	"crypto/rsa"
	"fmt"
)

var (
	ID = 0
)

var mappool = make([]Transaction, 0)

type Transaction struct {
	ID        int
	Operation *Operation
	nonce     uint
	FullSign  []byte
	PublicKey *rsa.PublicKey
	numOfSing int8
}

func (tx *Transaction) AddTX(o Operation) {

	if &o != nil {
		tx.Operation = &o
		tx.AddToMappool()
	}
}

func (tx *Transaction) ToString() string {
	return fmt.Sprintf("%d %s %d", tx.ID, tx.Operation.ToString(), tx.nonce)
}

func (tx *Transaction) AddToMappool() {
	mappool = append(mappool, *tx)
}
