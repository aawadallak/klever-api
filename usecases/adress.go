package usecases

import (
	"awesomeProject/domain"
	"awesomeProject/dto"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAdressDetails(address string) (*domain.Address, error){
	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/address/%s", address)
	req, err := utils.DoRequest(url, http.MethodGet, nil)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var body dto.AddressApiResponse
	err = json.Unmarshal(b, &body)
	if err != nil {
		return nil, err
	}

	balance := domain.NewBalance(body.Balance, body.UnconfirmedBalance)
	total := domain.NewTotal(body.TotalSent, body.TotalReceived)

	return domain.NewAddress(body.Address, body.Txs, balance, total), nil
}
