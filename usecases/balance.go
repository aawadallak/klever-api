package usecases

import (
	"awesomeProject/dto"
	"awesomeProject/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

func GetBalance(adresss string) (*dto.Balance, error){
	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", adresss)
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

	var confirmed big.Int
	var unconfirmed big.Int

	for _, v := range body {
		val, ok := new(big.Int).SetString(v.Value, 10)
		if !ok {
			return nil, errors.New("cannot conver bigint")
		}

		if v.Confirmations > 2 {
			confirmed.Add(&confirmed, val)
			continue
		}

		unconfirmed.Add(&unconfirmed, val)
	}

	return &dto.Balance{
		Confirmed:   confirmed.String(),
		Unconfirmed: unconfirmed.String(),
	}, nil
}