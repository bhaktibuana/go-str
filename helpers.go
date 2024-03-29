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

	if useDecimal {
		amountSplit := strings.Split(amountStr, ".")
		amountSplit[1] = amountSplit[1][:2]
		amountStr = strings.Join(amountSplit, decimal)
	}

	if dotSeparator && useDecimal {
		parts = []byte(strings.Join(strings.Split(amountStr, "."), ","))
	} else {
		parts = []byte(amountStr)
	}

	for i := range parts {
		if i > 0 && i%3 == 0 {
			formattedAmount = separator + formattedAmount
		}

		formattedAmount = string(parts[len(parts)-1-i]) + formattedAmount
	}

	if useDecimal {
		result := strings.Split(formattedAmount, decimal)

		index := strings.LastIndex(result[0], separator)
		if index == -1 {
			return strings.Join(result, decimal)
		}

		result[0] = result[0][:index]

		return strings.Join(result, decimal)
	}

	return formattedAmount
}
