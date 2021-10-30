package controllers

import (
	"awesomeProject/dto"
	"awesomeProject/usecases"
	"github.com/gin-gonic/gin"
)


func Send(c *gin.Context) {
	var bodyReq dto.SendRequest

	if err := c.ShouldBind(&bodyReq); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	v, err := usecases.SendTransaction(bodyReq.Amount, bodyReq.Address)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}
