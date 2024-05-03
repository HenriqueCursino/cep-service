package main

import (
	"cep-service/api"
	"cep-service/config"
	_ "cep-service/docs"
)

// @title Cep Sevice
// @version 1.0

// @description Cep Sevice
// @BasePath /
func main() {
	config.SetupConfigs()
	api.SetupApi()
}
