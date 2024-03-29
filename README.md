# Go String (goStr)

## Introduction

`goStr` is a Go package that provides functions for string manipulation, you are free to use them in your own applications if you find them convenient.

## Installation

To use `goStr` in your Go project, you can install it using the `go get` command:

```bash
go get github.com/bhaktibuana/go-str

```

## Available Functions

- [After](#after)
- [AfterLast](#afterlast)
- [APA](#apa)
- [Ascii](#ascii)
- [Before](#before)
- [BeforeLast](#beforelast)
- [Camel](#camel)
- [Headline](#headline)
- [Limit](#limit)
- [Snake](#snake)
- [Ucfirst](#ucfirst)
- [Ucsplit](#ucsplit)

### Examples

#### After

`After` method returns everything after the given value in a string. The entire string will be returned if the value does not exist within the string.

```go
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func After(input, substr string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.After("The quick brown fox jumps over the lazy dog", "brown")

// " fox jumps over the lazy dog"

```

#### AfterLast

`AfterLast` method returns everything after the last occurrence of the given value in a string. The entire string will be returned if the value does not exist within the string.

```go
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func AfterLast(input, substr string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.AfterLast("https://www.example.com/user/profile", "/")

// "profile"

```

#### APA

`APA` method converts the given string to title case following the [APA guidelines](https://apastyle.apa.org/style-grammar-guidelines/capitalization/title-case). (Supported language: English)

```go
/*
 * @param input string
 * @returns string
 */
func APA(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.APA("createing a project")

// "Creating a Project"

```

#### Ascii

`Ascii` method will attempt to transliterate the string into an ASCII value.

```go
/*
 * @param input string
 * @returns string
 */
func Ascii(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Ascii("û")

// "u"

```

#### Before

`Before` method returns everything before the given value in a string.

```go
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func Before(input, substr string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Before("The quick brown fox jumps over the lazy dog", "fox")

// "The quick brown "

```

#### BeforeLast

`BeforeLast` method returns everything before the last occurrence of the given value in a string.

```go
/*
 * @param input string
 * @param substr string
 * @returns string
 */
func BeforeLast(input, substr string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.BeforeLast("www.example.com/user/profile", "/")

// "www.example.com/user"

```

#### Camel

`Camel` method converts the given string to camelCase.

```go
/*
 * @param input string
 * @returns string
 */
func Camel(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Camel("foo_bar")

// "fooBar"

```

#### Headline

`Headline` method will convert strings delimited by casing, hyphens, or underscores into a space delimited string with each word's first letter capitalized.

```go
/*
 * @param input string
 * @returns string
 */
func Headline(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Headline("HelloWorld")

// "Hello World"

goStr.Headline("foo_bar")

// "Foo Bar"

```

#### Limit

`Limit` method truncates the given string to the specified length.

```go
/*
 * @param input string
 * @param length int
 * @param appendStr string
 * @returns string
 */
func Limit(input string, length int, appendStr ...string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Limit("The quick brown fox jumps over the lazy dog", 20)

// The quick brown fox...

```

You may pass a third argument to the method to change the string that will be appended to the end of the truncated string

```go
import "github.com/bhaktibuana/go-str"

goStr.Limit("The quick brown fox jumps over the lazy dog", 20, " (...)")

// The quick brown fox (...)

```

#### Snake

`Snake` method converts the given string to snake_case.

```go
/*
 * @param input string
 * @returns string
 */
func Snake(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Snake("myName-is john")

// "my_name_is_john"

```

#### Ucfirst

`Ucfirst` method returns the given string with the first character capitalized.

```go
/*
 * @param input string
 * @returns string
 */
func Ucfirst(input string) string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Ucfirst("foo bar")

// "Foo bar"

```

#### Ucsplit

`Ucsplit` method splits the given string into a collection by uppercase characters.

```go
/*
 * @param input string
 * @returns []string
 */
func Ucsplit(input string) []string
```

- usage

```go
import "github.com/bhaktibuana/go-str"

goStr.Ucsplit("HelloWorld")

// ["Hello", "World"]

```
