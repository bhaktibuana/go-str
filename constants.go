package goStr

var apaMinorWords = map[string]bool{
	"a":     true,
	"an":    true,
	"the":   true,
	"and":   true,
	"but":   true,
	"or":    true,
	"for":   true,
	"nor":   true,
	"on":    true,
	"at":    true,
	"to":    true,
	"by":    true,
	"with":  true,
	"of":    true,
	"in":    true,
	"into":  true,
	"near":  true,
	"over":  true,
	"off":   true,
	"up":    true,
	"down":  true,
	"out":   true,
	"as":    true,
	"about": true,
	"from":  true,
	"via":   true,
}

type CurrencyCode string

const (
	CURRENCY_IDR CurrencyCode = "IDR"
	CURRENCY_USD CurrencyCode = "USD"
	CURRENCY_EUR CurrencyCode = "EUR"
)

var CurrencyFormats = map[CurrencyCode]string{
	CURRENCY_IDR: "Rp",
	CURRENCY_USD: "$",
	CURRENCY_EUR: "€",
}
