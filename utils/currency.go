package utils

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// FormatRupiah formats a decimal.Decimal as "Rp. <amount>"
func FormatRupiah(amount decimal.Decimal) string {
	return fmt.Sprintf("Rp. %s", amount.StringFixed(0))
}
