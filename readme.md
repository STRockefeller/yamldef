# yamldef

![GitHub License](https://img.shields.io/github/license/STRockefeller/yamldef)![GitHub Top Language](https://img.shields.io/github/languages/top/STRockefeller/yamldef)![GitHub Actions Status](https://img.shields.io/github/actions/workflow/status/STRockefeller/yamldef/super-linter.yml)[![Go Report Card](https://goreportcard.com/badge/github.com/STRockefeller/go-linq)](https://goreportcard.com/report/github.com/STRockefeller/yamldef)[![Coverage Status](https://coveralls.io/repos/github/STRockefeller/go-linq/badge.svg?branch=main)](https://coveralls.io/github/STRockefeller/yamldef?branch=main)

yamldef is a Go module designed to generate Go source code from YAML definitions. It dynamically creates Go structs and provides serialization and deserialization methods for these structs using YAML data. This tool leverages the power of reflection and code generation to simplify working with YAML in Go applications.

## Features

- Dynamically generates Go structs from YAML data.
- Automatically implements `MarshalYAML` and `UnmarshalYAML` methods for the generated structs.
- Supports nested structures and arrays in YAML.
- Utilizes `jen` for code generation and `yaml.v3` for YAML parsing.

## Dependencies

yamldef uses the following Go modules:

- `github.com/dave/jennifer v1.7.0`: A code generation tool for Go.
- `github.com/stoewer/go-strcase v1.3.0`: A library for converting strings to various case formats.
- `gopkg.in/yaml.v3 v3.0.1`: The YAML package for Go, supporting YAML 1.2.

## Installation

To use yamldef in your project, ensure you have Go installed and your workspace is properly set up. Then, add yamldef to your project's dependencies:

```bash
go get github.com/STRockefeller/yamldef
```

## Usage

To generate Go source code from a YAML definition, use the `GenerateSourceCode` function. This function requires the YAML data as a byte slice, the directory path where the generated file should be saved, the package name for the generated file, and the struct name to be used in the generated code.

Example:

```go
package main

import (
	"github.com/STRockefeller/yamldef"
	"io/ioutil"
	"log"
)

func main() {
	ymlData, err := ioutil.ReadFile("example.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	err = yamldef.GenerateSourceCode(ymlData, "./model/", "def", "Person")
	if err != nil {
		log.Fatalf("Error generating source code: %s", err)
	}
}
```

This will generate a Go file in the specified directory with a struct definition based on the YAML content.

## License

yamldef is open-sourced software licensed under the MIT license.
