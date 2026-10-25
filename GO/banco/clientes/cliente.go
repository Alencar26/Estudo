package clientes

import (
	"regexp"
	"strconv"
)

type CPF string

type Titular struct {
	Nome, Profissao, CPF CPF
}

func (c CPF) ValidaCPF() bool {
	cpf := string(c)
	// Verifica se o formato está correto
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}\-\d{2}$`)
	if !re.MatchString(cpf) {
		return false
	}

	// Remove os caracteres não numéricos
	cpf = regexp.MustCompile(`\D`).ReplaceAllString(cpf, "")

	// Verifica se todos os dígitos são iguais
	if allDigitsEqual(cpf) {
		return false
	}

	// Calcula os dígitos verificadores
	d1 := calculateDigit(cpf[:9], 10)
	d2 := calculateDigit(cpf[:9]+strconv.Itoa(d1), 11)

	// Verifica se os dígitos verificadores são válidos
	return cpf[9] == byte(d1+'0') && cpf[10] == byte(d2+'0')
}

func allDigitsEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func calculateDigit(cpf string, factor int) int {
	sum := 0
	for _, digit := range cpf {
		num, _ := strconv.Atoi(string(digit))
		sum += num * factor
		factor--
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
