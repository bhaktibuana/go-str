package goStr

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/rainycape/unidecode"
)

// Camel method converts the given string to camelCase.
/*
 * @param input string
 * @returns string
 */
func Camel(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i, word := range words {
		if i > 0 {
			words[i] = strings.Title(word)
		}
	}

	return strings.Join(words, "")
}

// Snake snake method converts the given string to snake_case.
/*
 * @param input string
 * @returns string
 */
func Snake(input string) string {
	var firstBuilder strings.Builder
	var builder strings.Builder
	var finalBuilder strings.Builder
	var foundAlphanumeric bool
	var lastWasUnderscore bool

	for _, c := range input {
		if foundAlphanumeric || unicode.IsLetter(c) || unicode.IsDigit(c) {
			firstBuilder.WriteRune(c)
			foundAlphanumeric = true
		}
	}

	parsedInput := firstBuilder.String()

	for i, c := range parsedInput {
		if unicode.IsUpper(c) {
			if i > 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(c))
		} else if unicode.IsLetter(c) || unicode.IsDigit(c) {
			builder.WriteRune(c)
		} else if c != '_' {
			builder.WriteRune('_')
		}
	}

	parsedBuilder := builder.String()

	for _, c := range parsedBuilder {
		if c == '_' {
			if !lastWasUnderscore {
				finalBuilder.WriteRune(c)
			}
			lastWasUnderscore = true
		} else {
			finalBuilder.WriteRune(c)
			lastWasUnderscore = false
		}
	}
	return finalBuilder.String()
}

// Ucfirst method returns the given string with the first character.
/*
 * @param input string
 * @returns string
 */
func Ucfirst(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(string(input[0])) + input[1:]
}

// Ucsplit method splits the given string into an array by uppercase characters.
/*
 * @param input string
 * @returns []string
 */
func Ucsplit(input string) []string {
	var substrings []string
	start := 0

	for i, c := range input {
		if unicode.IsUpper(c) && i > 0 {
			substrings = append(substrings, input[start:i])
			start = i
		}
	}

	if start < len(input) {
		substrings = append(substrings, input[start:])
	}

	return substrings
}

// Headline method will convert strings delimited by casing, hyphens, or underscores into a space delimited string with each word's first letter capitalized.
/*
 * @param input string
 * @returns string
 */
func Headline(input string) string {
	input = strings.ReplaceAll(input, "-", " ")
	input = strings.ReplaceAll(input, "_", " ")

	var words []string
	var currentWord string

	for _, r := range input {
		if unicode.IsUpper(r) && currentWord != "" {
			words = append(words, currentWord)
			currentWord = string(r)
		} else {
			currentWord += string(r)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, " ")
}

// Limit method truncates the given string to the specified length.
/*
 * @param input string
 * @param length int
 * @param appendStr string
 * @returns string
 */
func Limit(input string, length int, appendStr ...string) string {
	if len(input) <= length {
		return input
	}
	if len(appendStr) > 0 {
		return input[:length] + appendStr[0]
	}
	return input[:length] + "..."
}

// After method returns everything after the given value in a string. The entire string will be returned if the value does not exist within the string.
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func After(input, substr string) string {
	index := strings.Index(input, substr)

	if index == -1 {
		return ""
	}

	return input[index+len(substr):]
}

// AfterLast method returns everything after the last occurrence of the given value in a string. The entire string will be returned if the value does not exist within the string.
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func AfterLast(input, substr string) string {
	index := strings.LastIndex(input, substr)

	if index == -1 {
		return input
	}

	return input[index+len(substr):]
}

// APA method converts the given string to title case following the APA guidelines. (Supported language: English)
/*
 * @param input string
 * @returns string
 */
func APA(input string) string {
	if input == "" {
		return input
	}

	words := strings.Fields(input)

	words[0] = strings.Title(words[0])

	for i := 1; i < len(words); i++ {
		if _, ok := apaMinorWords[strings.ToLower(words[i])]; ok {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}

	return strings.Join(words, " ")
}

// Ascii method will attempt to transliterate the string into an ASCII value.
/*
 * @param input string
 * @returns string
 */
func Ascii(input string) string {
	return unidecode.Unidecode(input)
}

// Before method returns everything before the given value in a string.
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func Before(input, substr string) string {
	index := strings.Index(input, substr)

	if index == -1 {
		return ""
	}

	return input[:index]
}

// BeforeLast method returns everything before the last occurrence of the given value in a string.
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func BeforeLast(input, substr string) string {
	index := strings.LastIndex(input, substr)

	if index == -1 {
		return input
	}

	return input[:index]
}

// Currency method will convert an amount of number or string into selected country code currency format.
/*
 * @param amount interface{} (int | float64 | string)
 * @param code CurrencyCode (string)
 * @param options ...bool (useDecimal true, dotSeparator false, useSpacer true)
 * @returns string
 */
func Currency(amount interface{}, code CurrencyCode, options ...bool) string {
	currencyFormat, ok := CurrencyFormats[code]
	if !ok {
		return ""
	}

	useDecimal := true
	if len(options) > 0 {
		useDecimal = options[0]
	}

	dotSeparator := false
	if len(options) > 0 {
		dotSeparator = options[1]
	}

	useSpacer := true
	if len(options) > 0 {
		useSpacer = options[2]
	}

	var num float64

	switch val := amount.(type) {
	case int:
		num = float64(val)
	case float64:
		num = val
	case string:
		parsedNum, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return ""
		}
		num = parsedNum
	default:
		return ""
	}

	amountStr := fmt.Sprintf("%f", num)
	if !useDecimal {
		amountStr = strconv.Itoa(int(num))
	}
	formattedAmount := FormatCurrency(amountStr, useDecimal, dotSeparator)

	if useSpacer {
		return fmt.Sprintf("%s %s", currencyFormat, formattedAmount)
	}

	return fmt.Sprintf("%s%s", currencyFormat, formattedAmount)
}
