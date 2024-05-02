package utils

import (
	"regexp"
	"strings"
)

func ValidateCEP(cep string) bool {
	regex := regexp.MustCompile(`\D`)
	unmaskedCep := regex.ReplaceAllString(cep, "")

	if len(unmaskedCep) != 8 {
		return false
	}

	match, _ := regexp.MatchString(`^\d{8}$`, unmaskedCep)
	return match
}

func ReplaceLastCepDigit(cep string) string {
	cepRune := []rune(cep)
	length := len(cepRune) - 1

	for i := length; i >= 0; i-- {
		if cepRune[i] != '0' {
			cepRune[i] = '0'
			break
		}
	}

	return string(cepRune)
}

func FormatCepUrl(url string, cep string) string {
	return strings.Replace(url, "?", cep, -1)
}
