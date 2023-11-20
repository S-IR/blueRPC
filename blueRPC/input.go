package bluerpc

import (
	"fmt"
	"os"
	"reflect"
)

func (a *App) Input(v interface{}) *App {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	file, err := os.Create(a.tsOutputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("interface %s {\n", t.Name()))
	if err != nil {
		panic(err)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tsType := "any"
		if field.Type.Kind() == reflect.String {
			tsType = "string"
		}

		_, err = file.WriteString(fmt.Sprintf("    %s: %s;\n", field.Name, tsType))
		if err != nil {
			panic(err)
		}
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		panic(err)
	}

	return a
}
