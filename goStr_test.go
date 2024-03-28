package goStr

import (
	"reflect"
	"testing"
)

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
		{"_1multiple__underscores_", "1multipleUnderscores"},
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

func TestSnake(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"FooBarBaz", "foo_bar_baz"},
		{"", ""},
		{"camelCase", "camel_case"},
		{"a-b-c", "a_b_c"},
		{"a b c", "a_b_c"},
		{"|my name-is johnDoe", "my_name_is_john_doe"},
		{"|1my name-is john", "1my_name_is_john"},
		{"I200 - Index Order Trucking", "i200_index_order_trucking"},
	}

	for _, tc := range testCases {
		result := Snake(tc.input)

		if result != tc.expected {
			t.Errorf("Snake(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestUcfirst(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello world"},
		{"this is my car", "This is my car"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := Ucfirst(tc.input)

		if result != tc.expected {
			t.Errorf("Ucfirst(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestUcsplit(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"HelloWorldFooBar", []string{"Hello", "World", "Foo", "Bar"}},
		{"camelCaseString", []string{"camel", "Case", "String"}},
		{"OneTwoThree", []string{"One", "Two", "Three"}},
		{"single", []string{"single"}},
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			result := Ucsplit(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Ucsplit(%q) = %q, expected %q", test.input, result, test.expected)
			}
		})
	}
}

func TestHeadline(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello_world", "Hello World"},
		{"FooBar", "Foo Bar"},
		{"john doe", "John Doe"},
		{"john-smith", "John Smith"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := Headline(tc.input)

		if result != tc.expected {
			t.Errorf("Headline(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}
