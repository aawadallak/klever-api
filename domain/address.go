package domain

type Address struct {
	Address string `json:"address"`
	TotalTx int `json:"totalTx"`
	Balance *Balance `json:"balance"`
	Total *Total `json:"total"`
}

func NewAddress(address string, totalTx int, balance *Balance, total *Total) *Address{
	return &Address{
		Address: address,
		TotalTx: totalTx,
		Balance: balance,
		Total:   total,
	}
}

