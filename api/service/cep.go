package service

import (
	"cep-service/api/response"
	"context"
	"errors"
)

type CepService interface {
	GetFirstAddress(cep string) (*response.GetAddressByCepResponse, error)
}

type cepService struct {
	urls map[string]func(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse)
}

func NewCepService(urls map[string]func(
	url string,
	cep string,
	ctx context.Context,
	responseChannel chan<- response.GetAddressByCepResponse,
)) CepService {
	return &cepService{urls: urls}
}

func (c *cepService) GetFirstAddress(cep string) (*response.GetAddressByCepResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	responseChannel := make(chan response.GetAddressByCepResponse)

	for url, callback := range c.urls {
		go func(
			ur string,
			cep string,
			callback func(string, string, context.Context, chan<- response.GetAddressByCepResponse),
			ctx context.Context) {
			callback(ur, cep, ctx, responseChannel)
		}(url, cep, callback, ctx)
	}

	select {
	case resp := <-responseChannel:
		if resp.Empty() {
			return nil, errors.New("response is empty")
		}

		return &resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
