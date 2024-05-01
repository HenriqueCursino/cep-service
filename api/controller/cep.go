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
		c.JSON(http.StatusBadRequest, responseFormatter.Error("CEP inválido"))
		return
	}

	ctx := c.Request.Context()
	response, err := cs.cepService.GetFirstAddress(cep, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, responseFormatter.Error(err))
		return
	}

	c.JSON(http.StatusOK, responseFormatter.Data(&response))
}
