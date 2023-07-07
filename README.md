# Legal Entity Identifier Go Library

A Go library for working with Legal Entity Identifiers (LEIs) as defined in [ISO 17442-1:2020](https://www.iso.org/standard/78829.html)

This library is a port of the [leim](https://gitlab.com/21analytics/lei) Rust library developed by 21 Analytics.

## Example

```go
import "github.com/trisacrypto/lei"

func main() {
    // Parse a valid LEI and check it
    if entityID, err := lei.Parse("2594007XIACKNMUAW223"); err != nil {
        log.Fatal(err)
    }
    fmt.Println(entityID.Check())

    // Parse an invalid LEI and print the invalid checksum error.
    _, err := lei.Parse("2594007XIACKNMUAW222")
    fmt.Println(err)
}
```

## Usage

Add the `lei` dependency to your `go.mod`:

```
$ go get github.com/trisacrypto/lei
```


## License

This project is licensed under the MIT license to match the license of the upstream Rust library this package was ported from.