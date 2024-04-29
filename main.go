package main

import (
	"cep-service/api"
	"cep-service/config"
)

func main() {
	config.SetupConfigs()
	api.SetupApi()
}
