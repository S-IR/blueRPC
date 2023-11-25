package genTypescript

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func GoFieldsToTSObj(someStruct reflect.Type) string {
	stringBuilder := strings.Builder{}

	if someStruct.Kind() != reflect.Struct {
		fmt.Printf("this is not a struct, it is a %s \n", someStruct.Kind())
		return "any"
	}

	stringBuilder.WriteString("{")

	for i := 0; i < someStruct.NumField(); i++ {
		field := someStruct.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Name()

		json := field.Tag.Get("json")
		input := field.Tag.Get("input")

		if input != "" {
			regex := regexp.MustCompile("[^a-zA-Z]+")
			fieldName = regex.ReplaceAllString(input, "")
		}

		if json != "" {
			regex := regexp.MustCompile("[^a-zA-Z]+")
			fieldName = regex.ReplaceAllString(json, "")
		}

		// Append TypeScript field definition to the StringBuilder
		stringBuilder.WriteString(fmt.Sprintf(" %s: %s", fieldName, goTypeToTSType(fieldType)))
		tags := field.Tag.Get("validate")
		if !strings.Contains(tags, "required") {
			stringBuilder.WriteString("|undefined")
		}
		stringBuilder.WriteString(",")

	}
	stringBuilder.WriteString("}")
	return stringBuilder.String()
}

func goTypeToTSType(goType string) string {
	switch goType {
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "number"
	case "float32", "float64":
		return "number"
	// Add more type mappings as needed
	default:
		return "any"
	}
}
