package dto

type AddressApiResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Txs int `json:"txs"`
	UnconfirmedBalance string `json:"unconfirmedBalance"`
	TotalSent string `json:"totalSent"`
	TotalReceived string `json:"totalReceived"`
}

type UTXOApiResponse []struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
}

type TransactionApiResponse struct {
	Block int `json:"block"`
	TxID string `json:"txID"`
	Vout []TransactionApiVOUT `json:"vout"`
}

type TransactionApiVOUT struct {
	BlockHeight int `json:"blockHeight"`
	Txid      string   `json:"txid"`
	Value     string   `json:"value"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}
