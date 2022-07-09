package Blockchain

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	FILE = "account_ID.txt"
)

type Account struct {
	ID        string
	Wallets   Keys
	Balance   int
	Validator bool
	Candidate bool
}

func (a *Account) CreateBlock(prevHash []byte) *Block {
	b := Block{
		"",
		prevHash,
		time.Now(),
		nil,
		a,
		nil,
	}

	id, err := genID()
	if err != nil {
		b.ID = "Empty ID"
	}
	b.ID = id
	b.AddTx()
	return &b
}

func (a Account) CreateOperation(receiver Account) (*Operation, error) {
	sender := a
	if receiver.Candidate {
		o := Operation{
			sender,
			receiver,
			1,
			nil,
		}

		sign, err := sender.SignData(o.ToString())
		if err != nil {
			return nil, err
		}
		o.Signature = sign
		fmt.Println("Op is created")
		o.CreateTxt()
		fmt.Println("Tx is created")
		return &o, nil
	}
	return nil, nil
}

func GenAccount() *Account {
	a := Account{}
	a.ID, _ = genID()
	file, err := os.OpenFile(FILE, os.O_APPEND, 0666)
	if err != nil {
		createFile, _ := os.Create(FILE)
		createFile.Close()

		file.Close()

		GenAccount()
	}
	data := make([]byte, 64)

	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	defer file.Close()
	allID := string(data)
	if !strings.Contains(allID, a.ID) {

		io.WriteString(file, fmt.Sprintln(a.ID))

		fmt.Println("Your ID is added")
	}
	a.Wallets = Keys{}
	//a.Wallets.PrivateKey, a.Wallets.PublicKey =
	a.Wallets.GenKeys()

	a.Validator = false
	a.Balance = 1
	a.Candidate = false
	return &a
}

func (a *Account) BecomeCandidate(c *Account) {
	if a.Validator {
		c.Candidate = true
	}
}

func (a *Account) UpdateBalance(balance int) {
	a.Balance = balance
}

func (a *Account) ChangeMyStatus() {
	a.Validator = true
}

func (a *Account) ChangeStatus(ac *Account) {
	if a.Validator || a != ac {
		ac.Validator = true
	}

}
func (a *Account) GetBalance() int {
	fmt.Println(a.Balance)
	return a.Balance
}

func (a *Account) VerifyTX(tx *Transaction) bool {
	if a.Validator {
		valid, _ := Verify(tx.PublicKey, tx.ToString(), tx.FullSign)
		if !valid {
			return false
		}
		return true
	}
	return false
}

func (a *Account) SignData(data string) ([]byte, error) {
	sign, err := a.Wallets.Sign(data, a.Wallets.GetPrivate())
	if sign == nil {
		return nil, err
	}
	return sign, nil
}

func (a *Account) ToString() string {
	priv, pub := a.Wallets.ToString()
	str := fmt.Sprintf("ID - %s\nBalance - %d\nPirvate - %s\nPublic - %s\n", a.ID, a.Balance, priv, pub)

	return str
}
