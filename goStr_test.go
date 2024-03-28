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

func TestLimit(t *testing.T) {
	firstTestCases := []struct {
		input    string
		length   int
		expected string
	}{
		{"Hello, my name is John.", 17, "Hello, my name is..."},
		{"Hello World", 0, "..."},
		{"", 0, ""},
	}

	secondTestCases := []struct {
		input     string
		length    int
		appendStr string
		expected  string
	}{
		{"Hello, my name is John.", 17, " (...)", "Hello, my name is (...)"},
		{"Hello World", 0, "(...)", "(...)"},
		{"", 0, "(...)", ""},
	}

	for _, tc1 := range firstTestCases {
		result := Limit(tc1.input, tc1.length)

		if result != tc1.expected {
			t.Errorf("Limit(%s) = %s; want %s", tc1.input, result, tc1.expected)
		}
	}

	for _, tc2 := range secondTestCases {
		result := Limit(tc2.input, tc2.length, tc2.appendStr)

		if result != tc2.expected {
			t.Errorf("Limit(%s) = %s; want %s", tc2.input, result, tc2.expected)
		}
	}
}

func TestAfter(t *testing.T) {
	testCases := []struct {
		input    string
		substr   string
		expected string
	}{
		{"The quick brown fox jumps over the lazy dog", "brown", " fox jumps over the lazy dog"},
		{"Yellow green blue", "red", ""},
		{"", "", ""},
	}

	for _, tc := range testCases {
		result := After(tc.input, tc.substr)

		if result != tc.expected {
			t.Errorf("After(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}
