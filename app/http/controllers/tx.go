package controllers

import (
	"awesomeProject/usecases"
	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context)  {
	tx := c.Param("tx")

	v, err := usecases.ListTransactions(tx)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}