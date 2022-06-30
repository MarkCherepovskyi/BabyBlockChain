package Account

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	FILE = "account_ID.txt"
)

type Account struct {
	ID      string
	Wallets Keys
	Balance int
}

func GenAccount() *Account {
	a := Account{}
	fmt.Fscan(os.Stdin, &a.ID)
	if a.ID == "" {
		fmt.Println("You don't enter ID")
	}
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

func (a Account) SignData(data string) []byte {
	sign := a.Wallets.Sign(data, a.Wallets.PrivateKey)
	if sign == nil {
		log.Println("ERROR")
		return nil
	}
	return sign
}

func (a Account) ToString() string {
	priv, pub := a.Wallets.ToString()
	str := fmt.Sprintf("ID %s \nBalance %s\nPirvate %s\nPublic %s\n", a.ID, string(a.Balance), priv, pub)

	fmt.Println(str)

	return str
}
