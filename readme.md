# YAML to Go Struct Converter

This Go package, `yamldef`, provides a utility to convert YAML data into Go struct definitions. It aims to simplify the process of working with YAML data by generating Go struct definitions that can be used directly in Go programs for marshaling and unmarshaling YAML data.

## Features

- **Automatic Struct Generation**: Converts YAML data into Go struct definitions automatically.
- **Support for Nested Structures**: Handles nested YAML structures, generating nested Go structs as needed.
- **Custom Type Handling**: Generates appropriate Go types (`int`, `string`, `[]interface{}`, etc.) based on the YAML data types.
- **Marshaling and Unmarshaling Functions**: Includes methods for marshaling and unmarshaling the generated structs to and from YAML.

## Installation

To use `yamldef` in your project, you need to install it first. Run the following command in your terminal:

```bash
go get github.com/STRockefeller/yamldef
```

## Usage

Here's a simple example of how to use `yamldef` to convert YAML data into a Go struct definition:

```go
package main

import (
	"fmt"
	"github.com/yourusername/yamldef"
)

func main() {
	yml := []byte(`
name: John Doe
age: 34
addresses:
  - type: home
    city: Somewhere
  - type: work
    city: Anywhere
languages: ["English", "Spanish"]
`)

	structDef := yamldef.YamlToStructDefinition(yml)
	fmt.Println(string(structDef))
}
```

This will output the Go struct definition based on the provided YAML data.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
