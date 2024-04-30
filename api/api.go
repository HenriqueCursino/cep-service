package api

import (
	"cep-service/config/env"

	"github.com/gin-gonic/gin"
)

func SetupApi() {
	server := gin.Default()

	Router(server)

	server.Run(env.Port)
}
