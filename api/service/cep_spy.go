package service

import "cep-service/api/response"

type CepServiceSpy struct {
	CepService
	GetFirstAddressResponse response.GetAddressByCepResponse
	GetFirstAddressError    error
}

func (c CepServiceSpy) GetAdressByCep(cep string) (*response.GetAddressByCepResponse, error) {
	return &c.GetFirstAddressResponse, c.GetFirstAddressError
}
