package utils

import (
	"github.com/shopspring/decimal"
)

// CalculatePriceAndFinal computes final price based on discount,
// formats "Rp." for display.
func CalculatePriceAndFinal(
	basePrice decimal.Decimal,
	isDiscount bool,
	discountRate decimal.Decimal,
) (string, string) {
	finalPrice := basePrice

	if isDiscount && discountRate.GreaterThan(decimal.Zero) {
		oneHundred := decimal.NewFromInt(100)
		finalPrice = basePrice.Mul(oneHundred.Sub(discountRate)).Div(oneHundred)
	}

	priceStr := FormatRupiah(basePrice)
	finalPriceStr := FormatRupiah(finalPrice)
	return priceStr, finalPriceStr
}
