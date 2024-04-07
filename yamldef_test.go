package yamldef

import (
	"os"
	"testing"
)

func TestGenerateSourceCode(t *testing.T) {
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
	err := GenerateSourceCode(yml, "./model/", "def", "Person")
	if err != nil {
		t.Errorf("Error generating source code: %s", err)
	}
	outputFile := "./model/person.go"
	expectedContent := `package def

import yamlv3 "gopkg.in/yaml.v3"

type Person struct {
	Addresses []struct {
		City       string ` + "`yaml:\"city\"`" + `
		PostalCode int    ` + "`yaml:\"postal_code\"`" + `
		Type       string ` + "`yaml:\"type\"`" + `
	} ` + "`yaml:\"addresses\"`" + `
	Age       int ` + "`yaml:\"age\"`" + `
	Education struct {
		Degree     string ` + "`yaml:\"degree\"`" + `
		University string ` + "`yaml:\"university\"`" + `
		Year       int    ` + "`yaml:\"year\"`" + `
	} ` + "`yaml:\"education\"`" + `
	Hobbies   []string ` + "`yaml:\"hobbies\"`" + `
	Languages []string ` + "`yaml:\"languages\"`" + `
	Name      string   ` + "`yaml:\"name\"`" + `
}

func (g Person) MarshalYAML() (interface{}, error) {
	return yamlv3.Marshal(g)
}
func (g *Person) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshal(g)
}
`

	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read the output file: %s", err)
	}

	if string(content) != expectedContent {
		t.Errorf("Generated file content does not match the expected content.")
	}
	if err := os.Remove(outputFile); err != nil {
		t.Fatalf("Failed to remove the output file: %s", err)
	}
}
