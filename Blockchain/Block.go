package Blockchain

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	MAX_NUM_OF_TX = 1
)

type Block struct {
	ID               string
	PrevHash         []byte
	time             time.Time
	setOfTransaction []*Transaction
	Signer           *Account
	Sign             []byte
}

func (b *Block) AddTx() {
	for _, tx := range Mappool {
		if len(b.setOfTransaction) < MAX_NUM_OF_TX {
			if b.Signer.VerifyTX(tx) {
				b.setOfTransaction = append(b.setOfTransaction, tx)
				Mappool = Mappool[:len(Mappool)-1]
				tx.Operation.carryOutOp()
			}
		} else if len(b.setOfTransaction) == MAX_NUM_OF_TX {

			b.SignBlock()
			return
		}
	}

}

func (b *Block) SignBlock() []byte {

	strBlock := b.ToString()
	sign, err := b.Signer.Wallets.Sign(strBlock, b.Signer.Wallets.GetPrivate())
	if err != nil {
		return nil
	}
	return sign
}

func (b *Block) Verify() bool {

	_, err := Verify(b.Signer.Wallets.GetPublicKey(), b.ToString(), b.Sign)
	if err != nil {
		return false
	}
	return true
}

func genID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("IDACC%s", id.String()), nil
}

func genIDBlock() (string, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("IDBLOCK%s", id.String()), nil

}
func (b *Block) ToString() string {
	return fmt.Sprintf("IDTX of Block - %s\nPrevHash - %s\nSetOfTx - %s\nTime - %s\n", b.ID, b.PrevHash, b.TxToString(), b.time.String())
}

func (b *Block) TxToString() string {
	buffer := ""
	for _, tx := range b.setOfTransaction {
		buffer += tx.ToString()
	}
	return buffer
}
