package controllers

import (
	"awesomeProject/usecases"
	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context){
	address := c.Param("address")

	v, err := usecases.GetBalance(address)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}