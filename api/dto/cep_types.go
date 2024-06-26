package dto

import (
	"cep-service/api/response"
)

type CepTypes interface {
	ViaCep | OpenCep | BrasilApi | BrasilAberto
}

// https://viacep.com.br/ws/01001000/json/
type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
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
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
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
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
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
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type BrasilAberto struct {
	Result ResultData `json:"result"`
}

type ResultData struct {
	Street         string `json:"street"`
	Complement     string `json:"complement"`
	District       string `json:"district"`
	City           string `json:"city"`
	State          string `json:"state"`
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
