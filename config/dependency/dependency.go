package dependency

import (
	"cep-service/api/controller"
	"cep-service/api/service"

	"github.com/rs/zerolog/log"
)

var (
	ViaCepController controller.CepController
)

func LoadDependencies() {
	log.Info().Msg("loading dependencies!")

	cepService := service.NewCepService()

	ViaCepController = controller.NewCepController(cepService)
}
