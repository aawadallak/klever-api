package controllers

import (
	"awesomeProject/usecases"
	"github.com/gin-gonic/gin"
)

func GetAdressDetails(c *gin.Context)  {
	address := c.Param("address")

	v, err := usecases.GetAdressDetails(address)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}

