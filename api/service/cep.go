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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	responseChannel := make(chan response.GetAddressByCepResponse)

	for url, callback := range c.urls {
		go func(url string, callback func(string, context.Context, chan<- response.GetAddressByCepResponse), ctx context.Context) {
			callback(utils.FormatCepUrl(url, cep), ctx, responseChannel)
		}(url, callback, ctx)
	}

	select {
	case resp := <-responseChannel:
		return &resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
