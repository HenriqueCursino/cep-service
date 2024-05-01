package api

import (
	"cep-service/config/dependency"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's running!",
		})
	})

	server.GET("cep/:cep", dependency.CepManagerController.GetAdressByCep)
}
