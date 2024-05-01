package dependency

import (
	"cep-service/api/controller"
	"cep-service/api/dto"
	"cep-service/api/response"
	"cep-service/api/service"
	"context"

	"github.com/rs/zerolog/log"
)

var (
	CepManagerController controller.CepController
)

func LoadDependencies() {
	log.Info().Msg("loading dependencies!")

	urls := map[string]func(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse){
		"https://viacep.com.br/ws/?/json/":          dto.GetViaCep,
		"https://opencep.com/v1/?":                  dto.GetOpenCep,
		"https://brasilapi.com.br/api/cep/v2/?":     dto.GetBrasilApi,
		"https://api.brasilaberto.com/v1/zipcode/?": dto.GetBrasilAberto,
	}

	cepService := service.NewCepService(urls)

	CepManagerController = controller.NewCepController(cepService)
}
