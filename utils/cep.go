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

func ReplaceLastCepDigit(input string) string {
	runes := []rune(input)
	length := len(runes) - 1

	for i := length; i >= 0; i-- {
		if runes[i] != '0' && runes[i] != '-' {
			runes[i] = '0'
			break
		}
	}

	return string(runes)
}

func HasNonZero(input string) bool {
	for _, char := range input {
		if char != '0' && char != '-' {
			return true
		}
	}
	return false
}

func FormatCepUrl(url string, cep string) string {
	return strings.Replace(url, "?", cep, -1)
}
