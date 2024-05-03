package controller

import (
	responseFormatter "cep-service/api/response"
	"cep-service/api/service"
	"cep-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CepController interface {
	GetAdressByCep(c *gin.Context)
}

type cepController struct {
	cepService service.CepService
}

func NewCepController(cepService service.CepService) CepController {
	return &cepController{
		cepService: cepService,
	}
}

func (cs *cepController) GetAdressByCep(c *gin.Context) {
	cep := c.Param("cep")
	if !utils.ValidateCEP(cep) {
		c.JSON(http.StatusBadRequest, responseFormatter.ResponseError{
			Message: "CEP inválido",
		})
		return
	}

	response, err := cs.cepService.GetFirstAddress(cep)
	if err != nil {
		if err.Error() == "response is empty" {
			c.JSON(http.StatusNotFound, responseFormatter.ResponseError{
				Message: "CEP inválido",
			})
			return
		}

		c.JSON(http.StatusBadRequest, responseFormatter.ResponseError{
			Message: "CEP inválido",
		})
		return
	}

	c.JSON(http.StatusOK, &response)
}
