package dto

import (
	"cep-service/api/response"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetViaCep(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[ViaCep](url, ctx).MapViaCepToResponse()
}
func GetOpenCep(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[OpenCep](url, ctx).MapOpenCepToResponse()
}

func GetBrasilApi(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[BrasilApi](url, ctx).MapBrasilApiToResponse()
}

func GetBrasilAberto(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
	responseChannel <- execute[BrasilAberto](url, ctx).MapBrasilAbertoToResponse()
}

func execute[T CepTypes](url string, ctx context.Context) T {
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
