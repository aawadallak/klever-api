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

func ListTransactions(tx string)  (*domain.Transaction, error) {
	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/tx/%s", tx)
	req, err := utils.DoRequest(url, http.MethodGet, nil)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	var body dto.TransactionApiResponse

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &body); err != nil {
		return nil, err
	}

	addresses := make(map[string]*big.Int, 0)

	for _, v := range body.Vout {
		amount, ok := new(big.Int).SetString(v.Value, 10)
		if !ok {
			return nil, errors.New("cannot convert to big.int")
		}

		for _, a := range v.Addresses {
			if addresses[a] == nil {
				addresses[a] = amount
				continue
			}

			addresses[a].Add(addresses[a], amount)
		}
	}

	response := domain.NewSendTransactionResponse(body.Block, body.TxID)

	for index, value := range addresses {
		response.Add(domain.NewAdressesInformation(index, value.String()))
	}

	return response,nil
}