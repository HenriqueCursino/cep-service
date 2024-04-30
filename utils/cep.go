package utils

import (
	"regexp"
	"strings"
)

func ValidateCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}

	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}

func FormatCepUrl(url string, cep string) string {
	return strings.Replace(url, "?", cep, -1)
}
