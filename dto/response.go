package dto

type BitcoinResponse struct {
	Address string `json:"address"`
	TotalTx int `json:"totalTx"`
	Balance Balance `json:"balance"`
	Total Total  `json:"total"`
}

type Total struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}

type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

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

type SendResponse struct {
	Utxos []SendResponseChild `json:"utxos"`
}

type SendResponseChild struct {
		Txid   string `json:"txid"`
		Amount string `json:"amount"`
}

type TransactionResponse struct {
	Block int `json:"block"`
	TxID string `json:"txID"`
	Vout []TransactionVOUT `json:"vout"`
}

type TransactionVOUT struct {
	BlockHeight int `json:"blockHeight"`
	Txid      string   `json:"txid"`
	Value     string   `json:"value"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type SendTransactionResponse struct {
	Addresses []SendAddressTransaction `json:"addresses"`
	Block int `json:"block"`
	TxID string `json:"txID"`
}

type SendAddressTransaction struct {
	Address string `json:"address"`
	Value string `json:"value"`
}