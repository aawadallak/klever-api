package controllers

import (
	"awesomeProject/dto"
	"awesomeProject/usecases"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/big"
	"net/http"
)

const (
	DEFAULT_TAX string = "120"
)

func Send(c *gin.Context) {
	tax, ok := new(big.Int).SetString(DEFAULT_TAX, 10)
	if !ok {
		c.JSON(400, gin.H{
			"error": "cannot convert to big.int",
		})
	}

	var bodyReq dto.SendRequest

	if err := c.ShouldBind(&bodyReq); err != nil {
		c.JSON(400, err)
	}

	amount, ok := new(big.Int).SetString(bodyReq.Amount, 10)
	if !ok {
		c.JSON(400, gin.H{
			"error": "cannot convert to big.int",
		})
	}

	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", bodyReq.Address)
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

	var body dto.UTXOApiResponse
	err = json.Unmarshal(b, &body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	var response []dto.SendResponseChild

	var sum *big.Int = new(big.Int).SetInt64(0)
	for _, v := range body {

		if v.Confirmations < 2 {
			continue
		}

		val, ok := new(big.Int).SetString(v.Value, 10)
		if !ok {
			c.JSON(400, gin.H{
				"error": "cannot convert to big.int",
			})
		}

		sum = sum.Add(sum, val)

		response = append(response, dto.SendResponseChild{
			Txid:   v.Txid,
			Amount: v.Value,
		})

		if  new(big.Int).Add(amount, tax).Cmp(sum) == -1 || new(big.Int).Add(amount, tax).Cmp(sum) == 0 {
			break
		}
	}

	v, err := usecases.GetBalance(bodyReq.Address)
	if err != nil {
		c.JSON(400, err)
		return
	}

	confirmed, ok := new(big.Int).SetString(v.Confirmed, 10)
	if !ok {
		c.JSON(400, gin.H{
			"error": "cannot convert to big.int",
		})

		return
	}

	if new(big.Int).Add(amount, tax).Cmp(confirmed) != 1 {
		c.JSON(200, dto.SendResponse{
			response,
		})

		return
	}

	c.JSON(200, gin.H{
		"error": "cannot complete transaction",
	})

}
