package usecases

import (
	"awesomeProject/domain"
	"awesomeProject/dto"
	"awesomeProject/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

const (
	DEFAULT_TAX string = "120"
)

func SendTransaction(amountRe string, address string) (*domain.Utxos, error){
	tax, ok := new(big.Int).SetString(DEFAULT_TAX, 10)
	if !ok {
		return nil, errors.New("cannot convert to big.int")
	}

	amount, ok := new(big.Int).SetString(amountRe, 10)
	if !ok {
		return nil, errors.New("cannot convert to big.int")
	}

	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", address)
	req, err := utils.DoRequest(url, http.MethodGet, nil)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var body dto.UTXOApiResponse
	err = json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}

	response := domain.NewUtxos()

	var sum *big.Int = new(big.Int).SetInt64(0)
	for _, v := range body {

		if v.Confirmations < 2 {
			continue
		}

		val, ok := new(big.Int).SetString(v.Value, 10)
		if !ok {
			return nil, errors.New("cannot convert to big.int")
		}

		sum = sum.Add(sum, val)

		response.Add(domain.NewTransacationAmount(v.Txid, v.Value))

		if  new(big.Int).Add(amount, tax).Cmp(sum) == -1 || new(big.Int).Add(amount, tax).Cmp(sum) == 0 {
			break
		}
	}

	v, err := GetBalance(address)
	if err != nil {
		return nil, err
	}

	confirmed, ok := new(big.Int).SetString(v.Confirmed, 10)
	if !ok {
		return nil, errors.New("cannot convert to big.int")
	}

	if new(big.Int).Add(amount, tax).Cmp(confirmed) != 1 {
		return response, nil
	}

	return nil, errors.New("insufficient funds")
}
