package yamldef

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
	"gopkg.in/yaml.v3"
)

func GenerateSourceCode(yml []byte, dirPath string, packageName string, structName string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	f := jen.NewFile(packageName)

	var data interface{}
	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		return err
	}

	if dataMap, ok := data.(map[string]interface{}); ok {
		structFields := make([]jen.Code, 0)
		keys := sortMapKeys(dataMap)
		for _, key := range keys {
			value := dataMap[key]
			fieldType := getType(value, 1)
			fieldName := strcase.UpperCamelCase(key)
			structFields = append(structFields, jen.Id(fieldName).Id(fieldType).Tag(map[string]string{"yaml": key}))
		}
		f.Type().Id(structName).Struct(structFields...)
	}

	// Add MarshalYAML method
	f.Func().Params(jen.Id("g").Id(structName)).Id("MarshalYAML").Params().Params(jen.Interface(), jen.Error()).Block(
		jen.Return(jen.Qual("gopkg.in/yaml.v3", "Marshal").Call(jen.Id("g"))),
	)

	// Add UnmarshalYAML method
	f.Func().Params(jen.Id("g").Op("*").Id(structName)).Id("UnmarshalYAML").Params(jen.Id("unmarshal").Func().Params(jen.Interface()).Error()).Error().Block(
		jen.Return(jen.Id("unmarshal").Call(jen.Id("g"))),
	)

	outputFilePath := fmt.Sprintf("%s/%s.go", dirPath, strings.ToLower(structName))
	fileContent := []byte(fmt.Sprintf("%#v", f))
	if err := os.WriteFile(outputFilePath, fileContent, 0644); err != nil {
		return err
	}
	fmt.Printf("Generated code successfully written to %s\n", outputFilePath)
	return nil
}

// generateStruct 遞歸生成結構體的字段定義。
func generateStruct(buf *bytes.Buffer, data interface{}, indent int) {
	indentStr := "    " // 使用4個空格作為縮進，增強可讀性
	if mapData, isMap := data.(map[string]interface{}); isMap {
		keys := sortMapKeys(mapData) // 獲取並排序鍵名，以保持順序一致
		for _, k := range keys {
			v := mapData[k]
			fieldType := getType(v, indent+1)
			fieldName := strcase.UpperCamelCase(k)
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
