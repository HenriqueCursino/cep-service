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

// ShowAccount godoc
// @Summary      Get address by CEP
// @Description  Api to get address by cep. This Api try to get the CEP info fastest way
// @Description  If the CEP is invalid, it replaces the last digit until a valid value is found.
// @Description  if no CEP is found returns Not Found
// @Tags 				 Cep
// @Accept       json
// @Produce      json
// @Param        cep   path      string  true  "CEP"
// @Success      200  {object}  responseFormatter.GetAddressByCepResponse
// @Failure      404  {object}  responseFormatter.ResponseError
// @Failure      401  {object}  responseFormatter.ResponseError
// @Failure      400  {object}  responseFormatter.ResponseError
// @Router       /cep/{cep} [get]
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
