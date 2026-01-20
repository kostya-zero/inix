# inix

**Inix** is a lightweight and strict INI parser for Go.

This parser disallows this:

- Inline comments are not allowed.
- Name of the section that contains spaces is not allowed.

## Installation

```shell
go get github.com/kostya-zero/inix
```

## Usage

```go
package main

import "github.com/kostya-zero/inix"

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
    value, ok := section.GetKey("data","hello")
    if !ok {
        // Error handling
    }

    fmt.Println(value)
}
```

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
