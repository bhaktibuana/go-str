package goStr

import "testing"

func TestCamel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello_world", "helloWorld"},
		{"foo bar baz", "fooBarBaz"},
		{"a_b_c", "aBC"},
		{"_leading|underscore", "leadingUnderscore"},
		{"trailing@underscore_", "trailingUnderscore"},
		{"_multiple__underscores_", "multipleUnderscores"},
		{"single", "single"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := Camel(tc.input)

		if result != tc.expected {
			t.Errorf("Camel(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}
