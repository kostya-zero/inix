# inix

**Inix** is a lightweight and strict INI parser for Go. 

> [!WARN]
> This parser doesnt not support inline comments! It will treat them as part of the value.

## Installation

```shell
go get github.com/kostya-zero/inix
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/kostya-zero/inix"
)

func main() {
    data := `[data]
hello=world
foo=bar`

    // Parse the data
    document, err := inix.Parse(data)
    if err != nil {
        // Error handling
    }

    // Get value of 'hello' key from a section 'data'
    value, ok := document.GetKey("data","hello")
    if !ok {
        // Error handling
    }

    fmt.Println(value)
}
```

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
