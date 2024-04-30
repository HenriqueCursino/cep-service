package config

import (
	"cep-service/config/dependency"
	"cep-service/config/env"
	"cep-service/config/log"
)

func SetupConfigs() {
	env.LoadEnvs()
	dependency.LoadDependencies()
	log.Load()
}
