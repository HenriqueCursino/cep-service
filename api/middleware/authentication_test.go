package middleware_test

import (
	"cep-service/api/middleware"
	"cep-service/config/env"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthenticationMiddleware(t *testing.T) {
	t.Run("Success - Authorization middleware should return 200 ok with valid token", func(t *testing.T) {
		env.Token = "token"

		router := gin.Default()
		w := httptest.NewRecorder()

		router.Use(middleware.AuthenticationMiddleware())

		router.GET("/teste_token", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Recurso protegido"})
		})

		validToken := "Bearer token"
		reqValidToken := httptest.NewRequest("GET", "/teste_token", nil)
		reqValidToken.Header.Set("Authorization", validToken)

		router.ServeHTTP(w, reqValidToken)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Fail - Authorization middleware in health route should return 200 ok with valid token", func(t *testing.T) {
		env.Token = "token_invalido"

		router := gin.Default()
		w := httptest.NewRecorder()

		router.Use(middleware.AuthenticationMiddleware())

		router.GET("/teste_token", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Recurso protegido"})
		})

		validToken := "Bearer token"
		reqValidToken := httptest.NewRequest("GET", "/teste_token", nil)
		reqValidToken.Header.Set("Authorization", validToken)

		router.ServeHTTP(w, reqValidToken)

		require.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
