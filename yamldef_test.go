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

	structDef := yamlToStructDefinition(yml)

	if trimSpacesForEachLine(string(structDef)) != trimSpacesForEachLine(expected) {
		t.Errorf("Generated struct definition does not match expected.\nExpected:\n%s\nGot:\n%s", trimSpacesForEachLine(expected), trimSpacesForEachLine(string(structDef)))
	}
}

func trimSpacesForEachLine(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
