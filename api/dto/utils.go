package dto

import (
	"cep-service/api/response"
	"cep-service/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func GetViaCep(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[ViaCep](url, cep, ctx).MapViaCepToResponse()
}

func GetOpenCep(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[OpenCep](url, cep, ctx).MapOpenCepToResponse()
}

func GetBrasilApi(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[BrasilApi](url, cep, ctx).MapBrasilApiToResponse()
}

func GetBrasilAberto(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[BrasilAberto](url, cep, ctx).MapBrasilAbertoToResponse()
}

func execute[T CepTypes](url string, cep string, ctx context.Context) T {
	newUrl := utils.FormatCepUrl(url, cep)
	fmt.Println(newUrl)

	var responseVar T
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, newUrl, nil)
	if err != nil {
		log.Err(err).Msgf("execute - Error to create request, url : %s", newUrl)
		return responseVar
	}

	client := http.Client{Timeout: time.Second * 1}
	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msgf("execute - Fail to get response, url : %s", newUrl)
		return responseVar
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return retryGetAdress[T](url, cep, ctx)
		}
		log.Info().Msgf("execute - Fail in response, url : %s - %d", newUrl, response.StatusCode)
		return responseVar
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Err(err).Msgf("execute - Error to read body, url : %s", newUrl)
		return responseVar
	}

	err = json.Unmarshal(body, &responseVar)
	if err != nil {
		log.Err(err).Msgf("execute - Error to unmarshal body, url : %s", newUrl)
		return responseVar
	}

	var empytStruct T
	if responseVar == empytStruct {
		return retryGetAdress[T](url, cep, ctx)
	}

	return responseVar
}

func retryGetAdress[T CepTypes](url string, cep string, ctx context.Context) T {
	if utils.HasNonZero(cep) {
		newCep := utils.ReplaceLastCepDigit(cep)
		return execute[T](url, newCep, ctx)
	}
	var empty T
	return empty
}
