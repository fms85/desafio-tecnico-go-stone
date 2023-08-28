package util

import (
	"github.com/shopspring/decimal"
)

func Sum(a float64, b float64) float64 {
	decimalA := ConvertToDecimal(a)
	decimalB := ConvertToDecimal(b)

	return decimalA.Add(decimalB).InexactFloat64()
}

func Sub(a float64, b float64) float64 {
	decimalA := ConvertToDecimal(a)
	decimalB := ConvertToDecimal(b)

	return decimalA.Sub(decimalB).InexactFloat64()
}

func ConvertToDecimal(number float64) decimal.Decimal {
	formattedNumber := decimal.NewFromFloat(number)

	return formattedNumber
}
