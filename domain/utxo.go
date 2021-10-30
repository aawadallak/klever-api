package domain

type Utxos struct {
	Utxos []*TransacationAmount `json:"utxos"`
}

func NewUtxos() *Utxos {
	return &Utxos{}
}

func NewUtxosWithTransactions(utxos []*TransacationAmount) *Utxos {
	return &Utxos{Utxos: utxos}
}

func (s *Utxos) Add(tx *TransacationAmount) {
	s.Utxos = append(s.Utxos, tx)
}

type TransacationAmount struct {
	Txid   string `json:"txid"`
	Amount string `json:"amount"`
}

func NewTransacationAmount(txid string, amount string) *TransacationAmount {
	return &TransacationAmount{Txid: txid, Amount: amount}
}
