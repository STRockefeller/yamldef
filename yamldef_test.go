package yamldef

import (
	"strings"
	"testing"
)

func TestYamlToStructDefinition(t *testing.T) {
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

	expected := `type GeneratedStruct struct {
		Addresses []struct {
			City string ` + "`yaml:\"city\"`" + `
			Type string ` + "`yaml:\"type\"`" + `
		} ` + "`yaml:\"addresses\"`" + `
		Age int ` + "`yaml:\"age\"`" + `
		Languages []string ` + "`yaml:\"languages\"`" + `
		Name string ` + "`yaml:\"name\"`" + `
}

func (g GeneratedStruct) MarshalYAML() (interface{}, error) {
    return yaml.Marshal(g)
}

func (g *GeneratedStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
    return unmarshal(g)
}
`

	structDef := YamlToStructDefinition(yml)

	if trimSpacesForEachLine(string(structDef)) != trimSpacesForEachLine(expected) {
		t.Errorf("Generated struct definition does not match expected.\nExpected:\n%s\nGot:\n%s", trimSpacesForEachLine(expected), trimSpacesForEachLine(string(structDef)))
	}
}

func TestYamlToStructDefinitionComplex(t *testing.T) {
	yml := []byte(`
name: Jane Smith
age: 29
addresses:
  - type: home
    city: Nowhere
    postal_code: 12345
  - type: work
    city: Everywhere
    postal_code: 67890
languages: ["French", "German", "Italian"]
hobbies: ["reading", "cycling", "hiking"]
education:
  university: State University
  degree: Bachelor's
  year: 2015
`)

	expected := `type GeneratedStruct struct {
		Addresses []struct {
			City string ` + "`yaml:\"city\"`" + `
			PostalCode int ` + "`yaml:\"postal_code\"`" + `
			Type string ` + "`yaml:\"type\"`" + `
		} ` + "`yaml:\"addresses\"`" + `
		Age int ` + "`yaml:\"age\"`" + `
		Education struct {
			Degree string ` + "`yaml:\"degree\"`" + `
			University string ` + "`yaml:\"university\"`" + `
			Year int ` + "`yaml:\"year\"`" + `
		} ` + "`yaml:\"education\"`" + `
		Hobbies []string ` + "`yaml:\"hobbies\"`" + `
		Languages []string ` + "`yaml:\"languages\"`" + `
		Name string ` + "`yaml:\"name\"`" + `
}

func (g GeneratedStruct) MarshalYAML() (interface{}, error) {
    return yaml.Marshal(g)
}

func (g *GeneratedStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
    return unmarshal(g)
}
`

	structDef := YamlToStructDefinition(yml)

	if trimSpacesForEachLine(string(structDef)) != trimSpacesForEachLine(expected) {
		t.Errorf("Generated struct definition does not match expected for complex case.\nExpected:\n%s\nGot:\n%s", trimSpacesForEachLine(expected), trimSpacesForEachLine(string(structDef)))
	}
}

func trimSpacesForEachLine(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
