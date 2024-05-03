package api

import (
	"cep-service/api/middleware"
	"cep-service/config/dependency"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's running!",
		})
	})
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("cep/:cep", middleware.AuthenticationMiddleware(), dependency.CepManagerController.GetAdressByCep)
}
