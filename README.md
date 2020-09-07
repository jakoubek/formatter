# formatter

**formatter** is a Go package with functions for formatting numbers.

## Installation

You can install this package in your project by running:

```
go get -u github.com/jakoubek/formatter
```

## Available functions

### Formatting an integer (int64)

```go
var number int64 = 12345
formattedNumber := formatter.FormatInteger(number)
fmt.Println(formattedNumber) // 12.345 
```

## Formatting a decimal value (float64)

```go
var number float64 = 12345.25
formattedNumber := formatter.FormatDecimal(number)
fmt.Println(formattedNumber) // 12.345,25
```
