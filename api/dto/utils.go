package dto

import (
	"cep-service/api/response"
	"cep-service/utils"
	"context"
	"encoding/json"
	"io"
	"net/http"

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
	var responseVar T
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Err(err).Msgf("execute - Error to create request, url : %s", url)
		return responseVar
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msgf("execute - Fail to get response, url : %s", url)
		return responseVar
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			retryGetAdress[T](url, cep, ctx)
		}
		log.Info().Msgf("execute - Fail in response, url : %s", url)
		return responseVar
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Err(err).Msgf("execute - Error to read body, url : %s", url)
		return responseVar
	}

	err = json.Unmarshal(body, &responseVar)
	if err != nil {
		log.Err(err).Msgf("execute - Error to unmarshal body, url : %s", url)
		return responseVar
	}

	return responseVar
}

func retryGetAdress[T CepTypes](url string, cep string, ctx context.Context) {
	newCep := utils.ReplaceLastCepDigit(cep)
	execute[T](utils.FormatCepUrl(url, newCep), newCep, ctx)
}
