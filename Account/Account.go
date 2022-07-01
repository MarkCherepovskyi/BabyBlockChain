package Account

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	FILE = "account_ID.txt"
)

type Account struct {
	ID      string
	Wallets Keys
	Balance int
	//Validator bool
}

func (a Account) CreateOperation(receiver Account, amount int) (*Operation, error) {
	sender := a
	o := Operation{
		sender,
		receiver,
		amount,
		nil,
	}

	sign, err := sender.SignData(o.ToString())
	if err != nil {
		return nil, err
	}
	o.Signature = sign

	return &o, nil

}

func (a Account) CreateTxt() *Transaction {
	tx := Transaction{
		ID,
		nil,
		0,
	}
	return &tx
}

func genID() (string, error) {
	ID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err

	}
	return string(ID), nil
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
	a.Wallets.PrivateKey, a.Wallets.PublicKey = a.Wallets.GenKeys()
	a.Balance = 1
	return &a
}

func (a Account) UpdateBalance(balance int) {
	a.Balance = balance

}

func (a Account) GetBalance() int {
	fmt.Println(a.Balance)
	return a.Balance
}

func (a Account) SignData(data string) ([]byte, error) {
	sign, err := a.Wallets.Sign(data, a.Wallets.PrivateKey)
	if sign == nil {

		return nil, err
	}
	return sign, nil
}

func (a Account) ToString() string {
	priv, pub := a.Wallets.ToString()
	str := fmt.Sprintf("ID %s \nBalance %d\nPirvate %s\nPublic %s\n", a.ID, a.Balance, priv, pub)

	return str
}
