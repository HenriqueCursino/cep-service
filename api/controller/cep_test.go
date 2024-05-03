package controller_test

import (
	"cep-service/api/controller"
	"cep-service/api/response"
	"cep-service/api/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCepController(t *testing.T) {
	t.Run("Success - GetAdressByCep should return 200 ok", func(t *testing.T) {
		expectedResponse := response.GetAddressByCepResponse{}
		gofakeit.Struct(&expectedResponse)

		cepService := service.CepServiceSpy{}

		fakeCep := "01001-000"

		controller := controller.NewCepController(cepService)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAdressByCep)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		if err != nil {
			return
		}

		response := response.GetAddressByCepResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}

		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, expectedResponse, response)
	})

	t.Run("Success - GetAdressByCep should return 400 badRequest", func(t *testing.T) {
		expectedResponse := response.GetAddressByCepResponse{}
		gofakeit.Struct(&expectedResponse)

		cepService := service.CepServiceSpy{}

		fakeCep := "0100100"

		controller := controller.NewCepController(cepService)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAdressByCep)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		if err != nil {
			return
		}

		response := response.GetAddressByCepResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}

		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Success - GetAdressByCep should return 404 notFound", func(t *testing.T) {
		expectedResponse := response.GetAddressByCepResponse{}
		gofakeit.Struct(&expectedResponse)

		cepService := service.CepServiceSpy{}

		fakeCep := "22288811"

		controller := controller.NewCepController(cepService)
		router := gin.Default()
		router.GET("/cep/:cep", controller.GetAdressByCep)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/cep/%s", fakeCep), nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		if err != nil {
			return
		}

		response := response.GetAddressByCepResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return
		}

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}
