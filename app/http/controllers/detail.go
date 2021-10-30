package controllers

import (
	"awesomeProject/dto"
	"awesomeProject/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAdressDetails(c *gin.Context)  {
	address := c.Param("address")

	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/address/%s", address)
	req, err := utils.DoRequest(url, http.MethodGet, nil)
	defer req.Body.Close()
	if err != nil {
		c.JSON(400, err)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	var body dto.AddressApiResponse
	err = json.Unmarshal(b, &body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, dto.BitcoinResponse{
		Address: body.Address,
		TotalTx: body.Txs,
		Balance: dto.Balance{
			Confirmed:   body.Balance,
			Unconfirmed: body.UnconfirmedBalance,
		},
		Total:   dto.Total{
			Sent:     body.TotalSent,
			Received: body.TotalReceived,
		},
	})
}

