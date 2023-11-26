package genTypescript

import (
	"fmt"
	"reflect"
	"strings"
)

func genTSFuncFromQuery(stringBuilder *strings.Builder, queryParams, output interface{}) {

	stringBuilder.WriteString("(")
	if queryParams != nil {
		qpType := reflect.TypeOf(queryParams)
		stringBuilder.WriteString("queryParams:")
		stringBuilder.WriteString(GoFieldsToTSObj(qpType))
	}
	stringBuilder.WriteString(")=>(")

	if output != nil {
		outputType := reflect.TypeOf(output)
		stringBuilder.WriteString(GoFieldsToTSObj(outputType))
	} else {
		stringBuilder.WriteString("void")
	}
	stringBuilder.WriteString(");")
}

func genTSFuncFromMutation(stringBuilder *strings.Builder, queryParams, input, output interface{}) {

	stringBuilder.WriteString("(")
	if queryParams != nil || input != nil {
		stringBuilder.WriteString("input : {")
	}

	if queryParams != nil {
		qpType := reflect.TypeOf(queryParams)
		if qpType.Kind() == reflect.Ptr {
			qpType = qpType.Elem()
		}
		stringBuilder.WriteString(fmt.Sprintf("queryParams:%s,", GoFieldsToTSObj(qpType)))
	}
	if input != nil {
		inputType := reflect.TypeOf(input)
		if inputType.Kind() == reflect.Ptr {
			inputType = inputType.Elem()
		}
		stringBuilder.WriteString(fmt.Sprintf("input:%s", GoFieldsToTSObj(inputType)))
	}

	if queryParams != nil || input != nil {
		stringBuilder.WriteString("}")
	}

	stringBuilder.WriteString(")=>(")
	if output != nil {
		outputType := reflect.TypeOf(output)
		if outputType.Kind() == reflect.Ptr {
			outputType = outputType.Elem()
		}
		stringBuilder.WriteString(GoFieldsToTSObj(outputType))
	} else {
		stringBuilder.WriteString("void")
	}
	stringBuilder.WriteString(");")
}
