# Go String (goStr)

## Introduction

`goStr` is a Go package that provides functions for string manipulation, you are free to use them in your own applications if you find them convenient.

## Installation

To use `goStr` in your Go project, you can install it using the `go get` command:

```bash
go get github.com/bhaktibuana/go-str

```

## Available Functions

- [Camel](#camel)
- [Headline](#headline)
- [Snake](#snake)
- [Ucfirst](#ucfirst)
- [Ucsplit](#ucsplit)

### Examples

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
