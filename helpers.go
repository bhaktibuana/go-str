package goStr

import (
	"strings"
)

// formatCurrency formats the currency amount with appropriate separators.
func FormatCurrency(amountStr string, useDecimal bool, dotSeparator bool) string {
	var formattedAmount string
	var parts []byte
	var separator string
	var decimal string

	if dotSeparator {
		separator = "."
		decimal = ","
	} else {
		separator = ","
		decimal = "."
	}

	if dotSeparator && useDecimal {
		parts = splitAmount(strings.Join(strings.Split(amountStr, "."), ","))
	} else {
		parts = splitAmount(amountStr)
	}

	for i := range parts {
		if i > 0 && i%3 == 0 {
			formattedAmount = separator + formattedAmount
		}

		formattedAmount = string(parts[len(parts)-1-i]) + formattedAmount
	}

	if useDecimal {
		result := strings.Split(formattedAmount, decimal)
		result[0] = BeforeLast(result[0], separator)

		return strings.Join(result, decimal)
	}

	return formattedAmount
}

// splitAmount splits the amount string into parts.
func splitAmount(amountStr string) []byte {
	return []byte(amountStr)
}
