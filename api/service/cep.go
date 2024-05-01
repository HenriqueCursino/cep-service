package service

import (
	"cep-service/api/dto"
	"cep-service/utils"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CepService interface {
	GetFirstAddress(cep string, ctx context.Context) (*dto.ViaCep, error)
}

type cepService struct {
}

func NewCepService() CepService {
	return &cepService{}
}

func (c *cepService) GetFirstAddress(cep string, ctx context.Context) (*dto.ViaCep, error) {
	response := make(chan dto.ViaCep)
	testUrl := utils.FormatCepUrl("https://viacep.com.br/ws/?/json/", cep)

	go c.sendRequest(testUrl, ctx, response)
	resp := <-response

	return &resp, nil
}

func (c *cepService) sendRequest(url string, ctx context.Context, responseChannel chan<- dto.ViaCep) {
	var finalResponse dto.ViaCep

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Err(err).Msgf("sendRequest - Error to create request")
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msgf("sendRequest - Fail to get response")
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Info().Msgf("sendRequest - Fail in response")
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Err(err).Msgf("sendRequest - Error to read body, url")
		return
	}

	err = json.Unmarshal(body, &finalResponse)
	if err != nil {
		log.Err(err).Msgf("sendRequest - Error to unmarshal body")
		return
	}

	responseChannel <- finalResponse
}
