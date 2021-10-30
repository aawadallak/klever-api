package controllers

import (
	"awesomeProject/dto"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/big"
	"net/http"
)

func GetTransaction(c *gin.Context)  {
	tx := c.Param("tx")

	url := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/tx/%s", tx)
	req, err := utils.DoRequest(url, http.MethodGet, nil)
	defer req.Body.Close()
	if err != nil {
		c.JSON(400, err)
		return
	}

	var body dto.TransactionResponse

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := json.Unmarshal(b, &body); err != nil {
		c.JSON(400, err)
		return
	}

	addresses := make(map[string]*big.Int, 0)

	for _, v := range body.Vout {
		amount, ok := new(big.Int).SetString(v.Value, 10)
		if !ok {
			c.JSON(400, gin.H{
				"error": "cannot convert to big.int",
			})
			return
		}

		for _, a := range v.Addresses {
			if addresses[a] == nil {
				addresses[a] = amount
				continue
			}

			addresses[a].Add(addresses[a], amount)
		}
	}

	var response dto.SendTransactionResponse
	var data []dto.SendAddressTransaction

	for index, value := range addresses {
		data = append(data, dto.SendAddressTransaction{
			Address: index,
			Value:   value.String(),
		})
	}

	response.TxID = body.TxID
	response.Addresses = data
	response.Block  = body.Block

	c.JSON(200, response)
}