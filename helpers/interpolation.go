package helpers

import (
	"strings"

	"github.com/shopspring/decimal"
)

type StringToDecimalInterceptor func(isEmpty bool, value *decimal.Decimal, err error)

// stringToDecimal
// Parameter : str string.
// Return : isEmpty bool, value decimal.Decimal, error error
func StringToDecimal(str string) (bool, decimal.Decimal, error) {
	str = strings.TrimSpace(str)
	if str == "" {
		return true, decimal.NewFromInt(0), nil
	}

	if !strings.Contains(str, ".") {
		str = str + ".00"
	}

	harga, err := decimal.NewFromString(str)
	if err != nil {
		return false, decimal.NewFromInt(0), err
	}

	return false, harga, nil
}

func StringTrimAndLower(str string) string {
	return strings.ToLower(strings.Trim(str, " "))
}