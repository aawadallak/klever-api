package domain

type Transaction struct {
	Addresses []*AdressesInformation `json:"addresses"`
	Block int `json:"block"`
	TxID string `json:"txID"`
}

type AdressesInformation struct {
	Address string `json:"address"`
	Value string `json:"value"`
}

func NewAdressesInformation(address string, value string) *AdressesInformation {
	return &AdressesInformation{Address: address, Value: value}
}

func NewSendTransactionResponse(block int, txID string) *Transaction {
	return &Transaction{Block: block, TxID: txID}
}

func (s *Transaction) Add(addressInfo *AdressesInformation) {
	s.Addresses = append(s.Addresses, addressInfo)
}



