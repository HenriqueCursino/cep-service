package dto

import (
	"cep-service/api/response"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CepTypes interface {
	ViaCep | OpenCep | BrasilApi | BrasilAberto
}

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

// https://viacep.com.br/ws/01001000/json/
type ViaCep struct {
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

func (v ViaCep) MapViaCepToResponse() response.GetAddressByCepResponse {
	return response.GetAddressByCepResponse{
		Street:       v.Logradouro,
		Neighborhood: v.Bairro,
		City:         v.Localidade,
		State:        v.Uf,
	}
}

// https://opencep.com/v1/15050305
type OpenCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro" type:"Street"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro" type:"Neighborhood"`
	Localidade  string `json:"localidade" type:"City"`
	Uf          string `json:"uf" type:"State"`
	Ibge        string `json:"ibge"`
}

func (o OpenCep) MapOpenCepToResponse() response.GetAddressByCepResponse {
	return response.GetAddressByCepResponse{
		Street:       o.Logradouro,
		Neighborhood: o.Bairro,
		City:         o.Localidade,
		State:        o.Uf,
	}
}

// https://brasilapi.com.br/api/cep/v2/01001000
type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state" type:"State"`
	City         string `json:"city" type:"City"`
	Neighborhood string `json:"neighborhood" type:"Neighborhood"`
	Street       string `json:"street" type:"Street"`
	Service      string `json:"service"`
}

func (b BrasilApi) MapBrasilApiToResponse() response.GetAddressByCepResponse {
	return response.GetAddressByCepResponse{
		Street:       b.Street,
		Neighborhood: b.Neighborhood,
		City:         b.City,
		State:        b.State,
	}
}

// https://cdn.apicep.com/file/apicep/06233-030.json
type ApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state" type:"State"`
	City       string `json:"city" type:"City"`
	District   string `json:"district" type:"Neighborhood"`
	Address    string `json:"address" type:"Street"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type BrasilAberto struct {
	Result ResultData `json:"result"`
}

type ResultData struct {
	Street         string `json:"street" type:"Street"`
	Complement     string `json:"complement"`
	District       string `json:"district" type:"Neighborhood"`
	City           string `json:"city" type:"City"`
	State          string `json:"state" type:"State"`
	StateShortname string `json:"stateShortname"`
	Zipcode        string `json:"zipcode"`
}

func (a BrasilAberto) MapBrasilAbertoToResponse() response.GetAddressByCepResponse {
	return response.GetAddressByCepResponse{
		Street:       a.Result.Street,
		Neighborhood: a.Result.District,
		City:         a.Result.City,
		State:        a.Result.State,
	}
}
