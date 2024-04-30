package service

import (
	"cep-service/utils"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CepService interface {
	GetFirstAddress(cep string, ctx context.Context) (*Viacep, error)
}

type cepService struct {
}

func NewCepService() CepService {
	return &cepService{}
}

type Viacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro" type:"Street"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro" type:"Neighborhood"`
	Localidade  string `json:"localidade" type:"City"`
	Uf          string `json:"uf" type:"State"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c *cepService) GetFirstAddress(cep string, ctx context.Context) (*Viacep, error) {
	testResponse := Viacep{}
	testUrl := utils.FormatCepUrl("https://viacep.com.br/ws/?/json/", cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, testUrl, nil)
	if err != nil {
		log.Err(err).Msgf("GetFirstAddress - Error to create request")
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msgf("GetFirstAddress - Fail to get response")
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Err(err).Msgf("GetFirstAddress - Error to read body, url")
		return nil, err
	}

	err = json.Unmarshal(body, &testResponse)
	if err != nil {
		log.Err(err).Msgf("GetFirstAddress - Error to unmarshal body")
		return nil, err
	}

	return &testResponse, err
}
