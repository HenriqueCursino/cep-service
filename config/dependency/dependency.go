package dependency

import (
	"cep-service/api/controller"
	"cep-service/api/service"

	"github.com/rs/zerolog/log"
)

var (
	CepManagerController controller.CepController
)

func LoadDependencies() {
	log.Info().Msg("loading dependencies!")

	cepService := service.NewCepService()

	CepManagerController = controller.NewCepController(cepService)
}
