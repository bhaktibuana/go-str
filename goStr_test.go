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

func TestAfterLast(t *testing.T) {
	testCases := []struct {
		input    string
		substr   string
		expected string
	}{
		{"https://www.example.com/user/profile", "/", "profile"},
		{"https://www.example.com/user/profile", "@", "https://www.example.com/user/profile"},
		{"", "", ""},
	}

	for _, tc := range testCases {
		result := AfterLast(tc.input, tc.substr)

		if result != tc.expected {
			t.Errorf("AfterLast(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestAPA(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Creating A Project", "Creating a Project"},
		{"creating a project", "Creating a Project"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := APA(tc.input)

		if result != tc.expected {
			t.Errorf("APA(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestAscii(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"û", "u"},
		{"ñ", "n"},
		{"é", "e"},
		{"ø", "o"},
		{"ß", "ss"},
		{"äöü", "aou"},
		{"こんにちは", "konnichiha"},
		{"你好世界", "Ni Hao Shi Jie "},
		{"안녕하세요", "annyeonghaseyo"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := Ascii(tc.input)

		if result != tc.expected {
			t.Errorf("Ascii(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestBefore(t *testing.T) {
	testCases := []struct {
		input    string
		substr   string
		expected string
	}{
		{"The quick brown fox jumps over the lazy dog", "fox", "The quick brown "},
		{"Yellow green blue", "red", ""},
		{"", "", ""},
	}

	for _, tc := range testCases {
		result := Before(tc.input, tc.substr)

		if result != tc.expected {
			t.Errorf("Before(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestBeforeLast(t *testing.T) {
	testCases := []struct {
		input    string
		substr   string
		expected string
	}{
		{"The quick brown fox jumps over the lazy dog", "fox", "The quick brown "},
		{"Yellow green blue", "red", "Yellow green blue"},
		{"www.example.com/user/profile", "/", "www.example.com/user"},
		{"", "", ""},
	}

	for _, tc := range testCases {
		result := BeforeLast(tc.input, tc.substr)

		if result != tc.expected {
			t.Errorf("BeforeLast(%s) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestCurrency(t *testing.T) {
	decimalTrue := true
	decimalFalse := false
	dotSeparatorTrue := true
	dotSeparatorFalse := false
	spacerTrue := true
	spacerFalse := false

	testCases := []struct {
		amount       interface{}
		code         CurrencyCode
		useDecimal   *bool
		dotSeparator *bool
		useSpacer    *bool
		expected     string
	}{
		{10000, CURRENCY_IDR, &decimalTrue, &dotSeparatorFalse, &spacerTrue, "Rp 10,000.00"},
		{"10000", CURRENCY_USD, &decimalTrue, &dotSeparatorFalse, &spacerTrue, "$ 10,000.00"},
		{12000, CURRENCY_IDR, &decimalTrue, &dotSeparatorTrue, &spacerTrue, "Rp 12.000,00"},
		{"12000", CURRENCY_IDR, &decimalFalse, &dotSeparatorFalse, &spacerTrue, "Rp 12,000"},
		{15250.8972, CURRENCY_IDR, &decimalTrue, &dotSeparatorTrue, &spacerFalse, "Rp15.250,89"},
		{"@", CURRENCY_IDR, &decimalTrue, &dotSeparatorTrue, &spacerFalse, ""},
		{"", CURRENCY_IDR, &decimalTrue, &dotSeparatorTrue, &spacerFalse, ""},
		{true, CURRENCY_IDR, &decimalTrue, &dotSeparatorTrue, &spacerFalse, ""},
		{1000, "", &decimalTrue, &dotSeparatorTrue, &spacerFalse, ""},
	}

	for _, tc := range testCases {
		var useDecimal bool
		if tc.useDecimal != nil {
			useDecimal = *tc.useDecimal
		} else {
			useDecimal = true
		}

		var dotSeparator bool
		if tc.dotSeparator != nil {
			dotSeparator = *tc.dotSeparator
		} else {
			dotSeparator = false
		}

		var useSpacer bool
		if tc.useSpacer != nil {
			useSpacer = *tc.useSpacer
		} else {
			useSpacer = true
		}

		result := Currency(tc.amount, tc.code, *tc.useDecimal, *tc.dotSeparator, *tc.useSpacer)

		if result != tc.expected {
			t.Errorf("Currency(%v, %s, %v, %v, %v) = %s; want %s", tc.amount, tc.code, useDecimal, dotSeparator, useSpacer, result, tc.expected)
		}
	}
}
