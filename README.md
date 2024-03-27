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
import "github.com/bhaktibuana/go-str"

goStr.Camel("foo_bar")

// "fooBar"

```

#### Snake

`Snake` converts a given string to snake_case.

```go
import "github.com/bhaktibuana/go-str"

goStr.Snake("myName-is john")

// "my_name_is_john"

```
