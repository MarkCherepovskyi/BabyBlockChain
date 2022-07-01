package Account

import "fmt"

var (
	ID = 0
)

type Transaction struct {
	ID              int
	setOfOperations []Operation
	nonce           uint
}

func (tx *Transaction) AddTX(o Operation) {

	tx.setOfOperations = append(tx.setOfOperations, o)

}

func (tx *Transaction) ToString() string {

	return fmt.Sprintf("%d %s %d", tx.ID, tx.opToString(), tx.nonce)

}

func (tx *Transaction) opToString() string {
	buffer := ""
	for _, data := range tx.setOfOperations {
		buffer += fmt.Sprintf("%s ", data.ToString())
	}
	return buffer
}
