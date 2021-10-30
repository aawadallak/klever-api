package server

import (
	"awesomeProject/app/http/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewSever() *Server {
	return &Server {
		port: "5000",
		server: gin.Default(),

	}
}

func (s *Server) Run() {

	routes := s.server.Group("api/v1/")
	{
		routes.GET("details/:address", controllers.GetAdressDetails)
		routes.GET("balance/:address", controllers.GetBalance)

		routes.POST("send", controllers.Send)
		routes.GET("tx/:tx", controllers.GetTransaction)

		routes.GET("health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"health": "good",
			})
		})
	}

	log.Fatal(s.server.Run(":"+s.port))
}
