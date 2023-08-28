package util

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var CpfValidator validator.Func = func(fl validator.FieldLevel) bool {
	cpf, ok := fl.Field().Interface().(string)
	if ok {
		if ValidateCPF(cpf) {
			return true
		}
	}

	return false
}

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := validateCPFSumDigit(firstPart, cpfFirstDigitTable)
	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)
	dsum := validateCPFSumDigit(secondPart, cpfSecondDigitTable)
	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)

	return finalPart == cpf
}

func validateCPFSumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0
	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}
