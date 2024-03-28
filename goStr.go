package goStr

import (
	"strings"
	"unicode"
)

// Camel converts a given string to camelCase.
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

// Snake converts a given string to snake_case.
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
