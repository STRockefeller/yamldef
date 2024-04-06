package yamldef

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// yamlToStructDefinition converts yaml file to go struct definition
func yamlToStructDefinition(yml []byte) []byte {
	var data interface{}
	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	structDef := bytes.Buffer{}
	structDef.WriteString("type GeneratedStruct struct {\n")
	generateStruct(&structDef, data, 1)
	structDef.WriteString("}\n\n")

	// marshal and unmarshal functions
	structDef.WriteString(`func (g GeneratedStruct) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(g)
}

func (g *GeneratedStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshal(g)
}
`)

	return structDef.Bytes()
}

// generateStruct 遞歸生成結構體的字段定義。
func generateStruct(buf *bytes.Buffer, data interface{}, indent int) {
	indentStr := "    " // 使用4個空格作為縮進，增強可讀性
	if mapData, isMap := data.(map[string]interface{}); isMap {
		keys := sortMapKeys(mapData) // 獲取並排序鍵名，以保持順序一致
		for _, k := range keys {
			v := mapData[k]
			fieldType := getType(v, indent+1)
			fieldName := cases.Title(language.English).String(k)
			buf.WriteString(fmt.Sprintf("%s%s %s `yaml:\"%s\"`\n", indentStr, fieldName, fieldType, k))
		}
	}
}

// getType returns the go type of the given yaml value
func getType(v interface{}, indent int) string {
	switch v := v.(type) {
	case int, int64, float64:
		return "int"
	case string:
		return "string"
	case []interface{}:
		if len(v) > 0 {
			firstElemType := getType(v[0], indent+1)
			if strings.HasPrefix(firstElemType, "struct") {
				// if the first element is a struct, then the whole array is a struct
				return fmt.Sprintf("[]%s", firstElemType)
			}
			return "[]" + firstElemType // if all elements are of the same type, then the whole array is of the same type
		}
		return "[]interface{}" // if the array is empty or the type is unknown, then the whole array is of type interface{}
	case map[string]interface{}:
		structDef := bytes.Buffer{}
		structDef.WriteString("struct {\n")
		generateStruct(&structDef, v, indent+1)
		structDef.WriteString(strings.Repeat("    ", indent) + "}") // make sure the struct ends with the correct indent
		return structDef.String()                                   // return the nested struct definition
	default:
		return "interface{}"
	}
}

// sortMapKeys returns the keys of the given map, sorted alphabetically
func sortMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
