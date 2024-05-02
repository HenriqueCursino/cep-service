package middleware

import (
	"cep-service/api/response"
	"cep-service/config/env"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != fmt.Sprintf("Bearer %s", env.Token) {
			ctx.JSON(http.StatusUnauthorized, response.Error(
				"Não autorizado",
			))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
