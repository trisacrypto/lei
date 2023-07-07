# Legal Entity Identifier Go Library

[![Go Reference](https://pkg.go.dev/badge/github.com/trisacrypto/lei.svg)](https://pkg.go.dev/github.com/trisacrypto/lei)
[![Tests](https://github.com/trisacrypto/lei/actions/workflows/test.yaml/badge.svg)](https://github.com/trisacrypto/lei/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/trisacrypto/lei)](https://goreportcard.com/report/github.com/trisacrypto/lei)

A Go library for working with Legal Entity Identifiers (LEIs) as defined in [ISO 17442-1:2020](https://www.iso.org/standard/78829.html)

This library is a port of the [leim](https://gitlab.com/21analytics/lei) Rust library developed by 21 Analytics.

## Example

```go
import "github.com/trisacrypto/lei"

func main() {
    // Check that a string is a valid LEI
    err := lei.LEI("2594007XIACKNMUAW223").Check()
    // err == nil

    // Parse invalid LEIs to return error information
    err = lei.LEI("2594007XIACKNMUAW222").Check()
    // errors.Is(err, lei.ErrInvalidChecksum) == true

    err = lei.LEI("25947XCKNMUAW223").Check()
    // errors.Is(err, lei.ErrInvalidLength) == true

    err = lei.LEI("2594007X#ACKNMUAW223").Check()
    // errors.Is(err, lei.ErrInvalidChar) == true

    // Check that a RegistrationAuthority is valid
    err = lei.CheckRA("RA777777")
    // err == nil

    // Lookup an unknown RegistrationAuthority
    err = lei.CheckRA("RA009099")
    // errors.Is(err, lei.ErrUnknownRA) == true
}
```

## Usage

Add the `lei` dependency to your `go.mod`:

```
$ go get github.com/trisacrypto/lei
```


## License

This project is licensed under the MIT license to match the license of the upstream Rust library this package was ported from.