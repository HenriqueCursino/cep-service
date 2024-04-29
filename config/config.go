package config

import (
	"cep-service/config/env"
	"cep-service/config/log"
)

func SetupConfigs() {
	env.LoadEnvs()
	log.Load()
}
