package utils

import (
	"time"

	"github.com/shopspring/decimal"
)

// CalculateBasePrice => base * (1 + additional)
func CalculateBasePrice(base decimal.Decimal, additional float64) decimal.Decimal {
	return base.Mul(decimal.NewFromFloat(1 + additional))
}

// CalculateFinalPrice => base * (1 - discount)
func CalculateFinalPrice(base decimal.Decimal, discount float64) decimal.Decimal {
	return base.Mul(decimal.NewFromFloat(1 - discount))
}

// CheckDiscountValid => helper validasi expired + aktif
func CheckDiscountValid(expired string, isActive *bool) bool {
	if isActive == nil || expired == "" {
		return false
	}
	exp, err := time.Parse(time.RFC3339, expired)
	if err != nil {
		return false
	}
	return *isActive && exp.After(time.Now())
}
