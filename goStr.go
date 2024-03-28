package goStr

import (
	"strings"
	"unicode"
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
