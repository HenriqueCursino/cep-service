package service

import (
	"cep-service/api/response"
	"context"
	"fmt"
	"log"
	"testing"

	"cep-service/api/dto/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	funcs := mock.NewMockCepFuncs(ctrl)

	// custom mocks
	funcs.EXPECT().GetViaCep(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
			responseChannel <- response.GetAddressByCepResponse{}
		})
	funcs.EXPECT().GetOpenCep(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	funcs.EXPECT().GetBrasilApi(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	funcs.EXPECT().GetBrasilAberto(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

	// mock element
	mocks := map[string]func(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse){
		"viacep":       funcs.GetViaCep,
		"opencep":      funcs.GetOpenCep,
		"brasilapi":    funcs.GetBrasilApi,
		"brasilaberto": funcs.GetBrasilAberto,
	}

	service := NewCepService(mocks)

	t.Run("should fail, reponse should be empty", func(t *testing.T) {
		resp, err := service.GetFirstAddress("14570000")

		log.Print("o que vieo 1", resp)

		assert.Equal(t, fmt.Errorf("response is empty"), err)
		assert.Nil(t, resp)
	})
}

func TestAdd2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	funcs := mock.NewMockCepFuncs(ctrl)

	// custom mocks
	funcs.EXPECT().GetViaCep(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse) {
			responseChannel <- response.GetAddressByCepResponse{City: "buritizão"}
		})

	// mock element
	mocks := map[string]func(url string, cep string, ctx context.Context, responseChannel chan<- response.GetAddressByCepResponse){
		"viacep": funcs.GetViaCep,
	}

	service := NewCepService(mocks)

	t.Run("meu teste de cep fail", func(t *testing.T) {
		resp, err := service.GetFirstAddress("14570000")

		log.Print("o que vieo 2", resp.City, err)

		assert.Nil(t, err)
		assert.Equal(t, "buritizão", resp.City)
	})
}
