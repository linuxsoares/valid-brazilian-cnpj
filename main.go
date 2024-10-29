package main

import (
	"strings"
)

func ValidCNPJ(cnpj string) bool {
	cnpj = onlyLetterNumber(cnpj)
	if len(cnpj) != 14 {
		return false
	}

	var sum int
	var weight int = 5
	for i := 0; i < 12; i++ {
		sum += int(cnpj[i]-'0') * weight
		weight--
		if weight == 1 {
			weight = 9
		}
	}

	remainder := sum % 11
	digit1 := 0
	if remainder < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - remainder
	}

	sum = 0
	weight = 6
	for i := 0; i < 13; i++ {
		sum += int(cnpj[i]-'0') * weight
		weight--
		if weight == 1 {
			weight = 9
		}
	}

	remainder = sum % 11
	digit2 := 0
	if remainder < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - remainder
	}

	return int(cnpj[12]-'0') == digit1 && int(cnpj[13]-'0') == digit2
}

func onlyLetterNumber(cnpj string) string {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	return cnpj
}
