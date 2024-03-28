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
- [Snake](#snake)

### Examples

#### Camel

`Camel` converts a given string to camelCase.

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

#### Snake

`Snake` converts a given string to snake_case.

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

`Ucfirst` returns the given string with the first character capitalized.

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

`Ucsplit` splits the given string into an array by uppercase characters.

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
