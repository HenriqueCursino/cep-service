package service

import (
	"cep-service/api/response"
	"cep-service/utils"
	"context"
)

type CepService interface {
	GetFirstAddress(cep string) (*response.GetAddressByCepResponse, error)
}

type cepService struct {
	urls map[string]func(url string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse)
}

func NewCepService(urls map[string]func(url string,
	ctx context.Context,
	responseChannel chan<- response.GetAddressByCepResponse)) CepService {
	return &cepService{urls: urls}
}

func (c *cepService) GetFirstAddress(cep string) (*response.GetAddressByCepResponse, error) {
	response := make(chan response.GetAddressByCepResponse)

	ctx := context.Background()
	defer ctx.Done()
	for url, callback := range c.urls {
		go callback(utils.FormatCepUrl(url, cep), ctx, response)
	}
	resp := <-response

	return &resp, nil
}
