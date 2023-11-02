package entity

type Transaction struct {
	Type  string
	Value float32
}

func NewTransaction(typeTransaction string, value float32) *Transaction {
	return &Transaction{
		Type:  typeTransaction,
		Value: value,
	}
}
