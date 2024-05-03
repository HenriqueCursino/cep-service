package utils_test

import (
	"cep-service/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateCEP(t *testing.T) {
	t.Run("Sucess - ValidateCEP", func(t *testing.T) {
		cepTest := "01001-000"
		expectedValue := true

		result := utils.ValidateCEP(cepTest)

		require.Equal(t, expectedValue, result)
	})

	t.Run("Fail - ValidateCEP cep length != 8", func(t *testing.T) {
		cepTest := "0100100"
		expectedValue := false

		result := utils.ValidateCEP(cepTest)

		require.Equal(t, expectedValue, result)
	})
}

func TestReplaceLastCepDigit(t *testing.T) {
	t.Run("Success - ReplaceLastCepDigit should change last digit to zero", func(t *testing.T) {
		expectedResult := "22223330"
		stringTest := "22223333"

		result := utils.ReplaceLastCepDigit(stringTest)

		require.Equal(t, expectedResult, result)
	})

	t.Run("Success - ReplaceLastCepDigit should return lasts digits to zero and not change '-'", func(t *testing.T) {
		expectedResult := "01000-000"
		stringTest := "01001-000"

		result := utils.ReplaceLastCepDigit(stringTest)

		require.Equal(t, expectedResult, result)
	})
}

func TestHasNonZero(t *testing.T) {
	t.Run("Sucess - HasNonZero should return false if dont have number diferent zero", func(t *testing.T) {
		cepTest := "00000-000"
		expectedValue := false

		result := utils.HasNonZero(cepTest)

		require.Equal(t, expectedValue, result)
	})

	t.Run("Sucess - HasNonZero should return true if have one number diferent zero", func(t *testing.T) {
		cepTest := "00001-000"
		expectedValue := true

		result := utils.HasNonZero(cepTest)

		require.Equal(t, expectedValue, result)
	})
}

func TestFormatCepUrl(t *testing.T) {
	t.Run("Sucess - FormatCepUrl should return formated url", func(t *testing.T) {
		urlTest := "https://viacep.com.br/ws/?/json/"
		cepTest := "01001-000"
		expectedValue := "https://viacep.com.br/ws/01001-000/json/"

		result := utils.FormatCepUrl(urlTest, cepTest)

		require.Equal(t, expectedValue, result)
	})
}
